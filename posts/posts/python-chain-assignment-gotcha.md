{
  "title": "Gotcha with Chained Assignment in Python",
  "status": "published",
  "slug": "python-chain-assignment-gotcha",
  "date": "2025-04-05T12:33:20Z"
}

<p>I was writing some Python code and wanted to initialize a few variables to an empty list. Instead of creating separate lists for each variable, I decided to use chained assignments like this:</p>
<pre><code class="language-python">a = b = c = [1,2,3]
</code></pre>
<p>It seems okay, nothing new, we are just assigning a, b, and c as an empty list. But notice the problem if you can.</p>
<p>If we try to change the value of b, what do you think is going to happen?</p>
<pre><code class="language-python">b.append(4)
</code></pre>
<p>If you didn’t know about how Python handles assignments, you might expect the variables to behave like this:</p>
<ul>
<li>
<p><code>a</code> should still be <code>[1, 2, 3]</code></p>
</li>
<li>
<p><code>b</code> should be <code>[1, 2, 3, 4]</code> (since I modified b)</p>
</li>
<li>
<p><code>c</code> should remain <code>[1, 2, 3]</code></p>
</li>
</ul>
<pre><code class="language-python"># You might think it will be this
# a = [1,2,3]
# b = [1,2,3,4]
# c = [1,2,3]
</code></pre>
<p>But, hello python, it hides the pointer magic behind this statement</p>
<p>The actual and expected state of the variables is:</p>
<pre><code class="language-python"># But it is actually this
# a = [1,2,3,4]
# b = [1,2,3,4]
# c = [1,2,3,4]
</code></pre>
<p>When you chain the assignment like this, all three variables refer to the same list object in memory. This means, that all the variables will hold the same object, this means that if you change any of the variables assigned that reference it will change the object, and that will result in changing all the variables since they are referring to the same variable.</p>
<p>You are not creating three independent lists. Instead, all three variables (a, b, and c) are referencing the same list object in memory. They don’t hold copies of the list, they all point to the same object. In Python, this is called <strong>reference assignment</strong>.</p>
<p>When you modify one of the variables (like appending to b), you’re not modifying just b, you’re modifying the single list object that all three variables are referencing. Since all three variables point to the same list, any change you make to the list will be reflected in all three variables.</p>
<p>On the other hand, if I do the same thing with strings, like this:</p>
<pre><code class="language-python">a = b = c = &quot;hello&quot;
b = &quot;world&quot;

# a = &quot;hello&quot;
# b = &quot;world&quot;
# c = &quot;hello&quot;
</code></pre>
<p>This only mutates the b variable and not the a and c, since string is not a mutable object in Python.</p>
<p>In Python, the objects are either <a href="https://realpython.com/python-mutable-vs-immutable-types/">mutable or immutable</a></p>
<p>Some of the primitive data types are:</p>
<p>Mutable Types:</p>
<ul>
<li>
<p>List</p>
</li>
<li>
<p>Dictionaries</p>
</li>
<li>
<p>Set</p>
</li>
<li>
<p>Byte Array</p>
</li>
</ul>
<p>Immutable Types:</p>
<ul>
<li>
<p>Integer, Float, Complex</p>
</li>
<li>
<p>String</p>
</li>
<li>
<p>Tuple</p>
</li>
<li>
<p>Bytes</p>
</li>
<li>
<p>Boolean</p>
</li>
<li>
<p>Frozenset</p>
</li>
</ul>
<p>So, suppose you assign multiple variables with the same value of a mutable type. In that case, the change in one variable will mutate the other variables as well since the underlying object in memory is the same.</p>
<p>So, this is what I learned from the mistake, avoid the chining assignment when dealing with mutable objects</p>
<p>Instead do the following:</p>
<pre><code class="language-python">a, b, c = [], [], []
b.append(256)

# a = []
# b = [256]
# c = []
</code></pre>
<p>This is safe and the right way to assign variables to individual values instead of the same value being referred to by all the variables.</p>
<p>Happy Coding :)</p>
