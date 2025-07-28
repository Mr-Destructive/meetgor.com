{
  "title": "Golang: Math Package",
  "status": "published",
  "slug": "golang-math",
  "date": "2025-04-05T12:33:42Z"
}

<h2>Introduction</h2>
<p>Moving on in the 100 days of golang series, we can take a look into the math package in golang's standard library. In programming, math is quite critical aspect, we need to perform certain mathematical operations quite regularly so golang's standard library has a package for serving some quite commonly used math functions and procedures. We'll take a look at some of the basic and common functions which are available in the math package.</p>
<h2>Mathematical Constants</h2>
<p>We have some constants like <code>pi</code>, <code>e</code>, <code>Phi</code> already defined as constants in the math package of the standard library in golang. They have a precision till 15 digits stored in float64 values.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;math&quot;
)

func main() {
	fmt.Println(&quot;Pi = &quot;, math.Pi)
	fmt.Println(&quot;E = &quot;, math.E)
	fmt.Println(&quot;Phi = &quot;, math.Phi)
	fmt.Println(&quot;Sqrt of 2 = &quot;, math.Sqrt2)
	fmt.Println(&quot;Naturla Log 2 = &quot;, math.Ln2)
	fmt.Println(&quot;Naturla Log 10 = &quot;, math.Ln10)
}
</code></pre>
<pre><code>$ go run basic-functions/constants.go
Pi =  3.141592653589793
E =  2.718281828459045
Phi =  1.618033988749895
Sqrt of 2 =  1.4142135623730951
Naturla Log 2 =  0.6931471805599453
Naturla Log 10 =  2.302585092994046
</code></pre>
<p>We can use these constants in trigonometric calculations and also in scientific computing. Further, you can get a list of all constants defined in the math package of the go standard library from the <a href="https://pkg.go.dev/math#pkg-constants">documentation</a>.</p>
<h2>Basic Math functions</h2>
<p>We have some quite basic and fundamental functions in the math package that can be used commonly in many programs. Let's take a look at a few of them.</p>
<h3>- Abs :parameters (float64) , returns float64</h3>
<p>As the name suggest, the <a href="https://pkg.go.dev/math#Abs">Abs</a> it returns the absolute result of a numbers. It takes a parameter as a float64 value and returns the absolute value of the provided number as a <code>float64</code> number.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;math&quot;
)

func main() {
	a := 45
	b := 100
	diff := a - b
	fmt.Println(diff)

	absolute_diff := math.Abs(float64(a) - float64(b))
	fmt.Println(absolute_diff)
}
</code></pre>
<pre><code>$ go run basic-functions/main.go
-55
55
</code></pre>
<p>As, we can see the <code>Abs</code> function takes in a float64 value and returns a absolute value of the given number that too a <code>float64</code> value. We need to cast the numbers <code>a</code> and <code>b</code> into <code>float64</code> as we have not provided the initial values and so the compiler has assigned the type to them as <code>int</code>.</p>
<h4>Type Casting</h4>
<p>We can caste a type into other by using the variable around the type name as <code>type_name(variable)</code>. In the above example we have converted the <code>int</code> value <code>45</code> into a <code>float64</code> as <code>float64(45)</code> which again yields <code>45</code> but as a float64 type.</p>
<pre><code class="language-go">foo := 77
fmt.Printf(&quot;Type of foo = %T 
&quot;, foo)
fmt.Println(&quot;foo = &quot;, int(foo))
fmt.Println(&quot;String Cast: &quot;, string(foo))
fmt.Println(&quot;Float Cast: &quot;, float64(foo))
</code></pre>
<p>Though not every type cannot be casted into due to quite oblivious reasons, for instance <code>77</code> or any other integer value (except for 0 or 1) cannot be converted into boolean value.</p>
<p><strong>Hello Gopher! Just a small note, the math package almost deals with float64 types rather than int to avoid backwards compatibility to perform operations on floating point values which can be casted into integers rather than defining separate functions for decimal values and integers.</strong></p>
<h3>- Min/Max: parameters(float64) , returns float64</h3>
<p>We can get the <a href="https://pkg.go.dev/math#Min">minimum</a> and <a href="https://pkg.go.dev/math#Max">maximum</a> value of the two numbers provided to the function.</p>
<pre><code class="language-go">var float64 a = 120
var float64 b = 54

minimum := math.Min(float64(a), float64(b))
maximum := math.Max(float64(a), float64(b))
fmt.Printf(&quot;Min of %v and %v is %v 
&quot;, a, b, minimum)
fmt.Printf(&quot;Max of %v and %v is %v 
&quot;, a, b, maximum)
</code></pre>
<pre><code>$ go run basic-functions/main.go
Min of 120 and 54 is 54
Max of 120 and 54 is 120
</code></pre>
<h3>- Pow  : parameters(float64, float64) , returns float64</h3>
<h3>- Pow10: parameters(int)              , returns float64</h3>
<p>The <a href="https://pkg.go.dev/math#Pow">Pow</a> function is used to get the exponential result of the base number. So, if we provide the values x and y, we would get the result as the number x raised to y.</p>
<pre><code class="language-go">var x float64 = 3
var y float64 = 4
z := math.Pow(x, y)
z10 := math.Pow10(int(x))
fmt.Println(&quot;X ^ Y = &quot;, z)
fmt.Println(&quot;10 ^ X = &quot;, z10)
</code></pre>
<pre><code>$ go run basic-functions/main.go
X ^ Y =  81
10 ^ X =  1000
</code></pre>
<p>We also have the <a href="https://pkg.go.dev/math#Pow10">Pow10</a> function which works just like the pow function except the x value is 10 and we don't have to provide it, there is just one parameter as a integer which returns a float64 value.</p>
<h3>- Sqrt: parameters(float64) , returns float64</h3>
<p>The <a href="https://pkg.go.dev/math#Sqrt">Sqrt</a> function as the name suggest, it is used to get the square root value of a floating point value which returns a float64 value.</p>
<pre><code class="language-go">var k float64 = 125
sqrt_of_k := math.Sqrt(k)
cbrt_of_k := math.Cbrt(k)

fmt.Printf(&quot;Square root of %v = %v 
&quot;, k, sqrt_of_k)
fmt.Printf(&quot;Cube root of %v = %v 
&quot;, k, cbrt_of_k)
</code></pre>
<pre><code>$ go run basic-functions/main.go
Square root of 125 = 11.180339887498949
Cube root of 125 = 5
</code></pre>
<h3>- Trunc: parameters(float64) , returns float64</h3>
<p>The <a href="https://pkg.go.dev/math#Trunc">Truncate</a> function provides the way to round off a decimal value(float64) to an integer but it returns a value in <code>float64</code>.</p>
<pre><code class="language-go">var p float64 = 445.235
trunc_p := math.Trunc(p)
fmt.Printf(&quot;Truncated value of %v = %v 
&quot;, p, trunc_p)
p = 123.678
trunc_p = math.Trunc(p)
fmt.Printf(&quot;Truncated value of %v = %v 
&quot;, p, trunc_p)

</code></pre>
<pre><code>$ go run basic-functions/main.go
Truncated value of 445.235 = 445
Truncated value of 123.678 = 123
</code></pre>
<h3>- Ceil : parameters(float64) , returns float64</h3>
<p>We also can use the <a href="https://pkg.go.dev/math#Ceil">Ceil</a> function to roud up the value to the next integer value but the value is returned as <code>float64</code>.</p>
<pre><code class="language-go">var c float64 = 33.25
ceil_c := math.Ceil(c)
fmt.Printf(&quot;Ceiled value of %v = %v 
&quot;, c, ceil_c)
c = 134.78
ceil_c = math.Ceil(c)
fmt.Printf(&quot;Ceiled value of %v = %v 
&quot;, c, ceil_c)
</code></pre>
<pre><code>$ go run basic-functions/main.go
Ceiled value of 33.25 = 34
Ceiled value of 134.78 = 135
</code></pre>
<h3>- Trigonometric Functions</h3>
<p>Trigonometric functions are quite helpful that can help in intense mathematical computations in backend projects or precision dependent projects. We have functions <a href="https://pkg.go.dev/math#Sin">Sin</a>, <a href="https://pkg.go.dev/math#Cos">Cos</a>, <a href="https://pkg.go.dev/math#Sincos">SinCos</a>, <a href="https://pkg.go.dev/math#Tan">Tan</a>, hyperbolic functions in Trigonometric functions like <a href="https://pkg.go.dev/math#Sinh">Sinh</a>, <a href="https://pkg.go.dev/math#Cosh">Cosh</a>, <a href="https://pkg.go.dev/math#Tanh">Tanh</a>, and Inverse Trigonometric functions like <a href="https://pkg.go.dev/math#Asin">Asin</a>, <a href="https://pkg.go.dev/math#Asinh">Asinh</a>, etc.</p>
<ul>
<li>Sin: parameters(float64) , returns float64</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;math&quot;
)

func main() {
	// basic trigonometric functions
	var x float64 = math.Pi / 2
	sinx := math.Sin(x)
	cosx := math.Cos(x)
	tanx := math.Tan(x)
	fmt.Printf(&quot;Sin(%v) = %v 
&quot;, x, sinx)
	fmt.Printf(&quot;Cos(%v) = %v 
&quot;, x, cosx)
	fmt.Printf(&quot;Tan(%v) = %v 
&quot;, x, tanx)

	// hyperbolic trigonometric functions
	var h float64 = math.Pi / 2
	sinh := math.Sinh(h)
	cosh := math.Cosh(h)
	tanh := math.Tanh(h)
	fmt.Printf(&quot;Sinh(%v) = %v 
&quot;, h, sinh)
	fmt.Printf(&quot;Cosh(%v) = %v 
&quot;, h, cosh)
	fmt.Printf(&quot;Tanh(%v) = %v 
&quot;, h, tanh)

	// Inverse Trigonometric functions
	var y float64 = -1
	arc_sin := math.Asin(y) // -pi/2 radians or 90 degrees
	arc_cos := math.Acos(y) // pi randians or 180 degrees
	arc_tan := math.Atan(y) 
	fmt.Printf(&quot;Sin^-1(%v) = %v 
&quot;, y, arc_sin)
	fmt.Printf(&quot;Cos^-1(%v) = %v 
&quot;, y, arc_cos)
	fmt.Printf(&quot;Tan^-1(%v) = %v 
&quot;, y, arc_tan)

</code></pre>
<pre><code>$ go run basic-functions/trignometric.go
Sin(1.5707963267948966) = 1
Cos(1.5707963267948966) = 6.123233995736757e-17
Tan(1.5707963267948966) = 1.6331239353195392e+16
Sinh(1.5707963267948966) = 2.3012989023072947
Cosh(1.5707963267948966) = 2.5091784786580567
Tanh(1.5707963267948966) = 0.9171523356672744
Sin^-1(-1) = -1.5707963267948966
Cos^-1(-1) = 3.141592653589793
Tan^-1(-1) = -0.7853981633974483
</code></pre>
<p>Here we can see that the functions are working fine and giving a decently precise value. This might be enough for simple and smaller projects, though for higher precision and accuracy areas, other computations and programming is required to compute the values.</p>
<h3>- Exponential and Logarithmic Functions</h3>
<p>We also have the exponential and logarithmic functions defined in the math package to leverage computations realted to formulae that deal with logarithmic or exponential calculations.</p>
<ul>
<li>Exp  : parameters(flaot64) , returns float64</li>
<li>Exp2 : parameters(flaot64) , returns float64</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;math&quot;
)

func main() {
	// exponential function
	var x float64 = 2
	y := math.Exp(x)
	fmt.Println(&quot;e^x = &quot;, y)
	var n float64 = 3.5
	y = math.Exp2(n)
	fmt.Println(&quot;2^n = &quot;, y)

	// Logarithmic function
	y = math.Log(x)
	fmt.Println(&quot;natural log x = &quot;, y)

	n = 128
	y = math.Log2(n)
	fmt.Println(&quot;Log2 of 100 = &quot;, y)
}
</code></pre>
<pre><code>$ go run basic-functions/expo_log.go
e^x =  7.38905609893065
2^n =  11.31370849898476
natural log x =  0.6931471805599453
Log2 of 100 =  7
</code></pre>
<p>Here, we have exponential functions such as <code>e^x</code> and <code>2^n</code> which might be useful in some common programming calculations. Also the logarithmic functions like <code>log x</code> which is natural log of x(base e), and <code>log2 n</code> which is logn to the base 2.</p>
<h2>The Random package</h2>
<p>The <code>random</code> sub-package in golang provides some great tools for working with random numbers and generating them. It provides exhaustive list of functions and types that help in generating pseudo random numbers.</p>
<ul>
<li>Int : parameters() , returns int</li>
<li>Intn : parameters( int ) , returns int</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;math/rand&quot;
)

func main() {
	// random integer generation
	x := rand.Int()
	fmt.Println(x)

	// random number generation till range
	for i := 0; i &lt; 5; i++ {
		y := rand.Intn(10)
		fmt.Println(y)
	}
}
</code></pre>
<pre><code>$ go run basic-functions/rand.go
5577006791947779410
7
7
9
1
8
</code></pre>
<p>In the above example, we have used the <a href="https://pkg.go.dev/math/rand@go1.18.1#Int">Int</a> function in the random sub-package of the math package which generates a pseudo random integer of the range dependent on the system architecture generally <code>int32</code> or <code>int64</code>. We get a huge number which is pseudo random i.e. not truly random. If you try to execute the program a couple of time, you would notice the number remains the same and we are calling it random? Well we need to dive into random numbers and seeding for a different part of the series for sure.</p>
<p>The <a href="https://pkg.go.dev/math/rand@go1.18.1#Intn">Intn</a> function also generates a pseudo random number but this time, we define the range of the upper boundary to generate them. It is not inclusive of the number provided i.e. we have provide the value <code>10</code> so the number <code>10</code> is not included in the range. It's called half open interval. It starts from 0 so the range becomes mathematically <code>[0, n)</code> if n is the number provided to the Intn function.</p>
<h2>The Bits package</h2>
<p>We also have a bit sub-package in the math package of the go standard library. This sub package is used for working around with bit manipulation and operations at the binary level. This is quite helpful in competitive programming , also in understanding the basics of data structures and fundamentals in computer science.</p>
<ul>
<li>Add       : parameters(uint, uint, uint) , returns uint, uint</li>
<li>Len       : parameters(uint) , returns int</li>
<li>OnesCount : parameters(uint) , returns int</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;math/bits&quot;
)

func main() {
	s, c := bits.Add(0, 9, 1)
	fmt.Printf(&quot;Sum = %d 
Carry = %d 
&quot;, s, c)

	// (45) in decimal = (1 0 1 1 0 1) in binary
	var n uint = 45
	length := bits.Len(n)
	ones_in_45 := bits.OnesCount(n)
	fmt.Printf(&quot;Minimum bits required to represent 45 = %d 
&quot;, length)
	fmt.Printf(&quot;Set Bits in 45 = %d 
&quot;, ones_in_45)

}
</code></pre>
<pre><code>$ go run basic-functions/bit.go
Sum = 10
Carry = 0
Minimum bits required to represent 45 = 6
Set Bits in 45 = 4
</code></pre>
<p>Here, in the above example, we have used the bits sub pacakge in the math package, the <a href="https://pkg.go.dev/math/bits@go1.18.1#Add">Add</a> function allows us to provide the two numbers and a carry bit on which it returns two values the sum and the carry. The sum is defined as the summation of <code>x + y + carry</code> the two numbers and the carry bit. The carry bit needs to be either 0 or 1.</p>
<p>Also the value provided the function i.e. <code>x and y</code> need to be unsigned <code>uint</code> iorder to work with bits.</p>
<p>We also have the <a href="https://pkg.go.dev/math/bits@go1.18.1#Len">Len</a> function which returns the maximum number of bits required to represent the provided unsigned integer. We have used 45 which is equivalent to <code>10110</code> and hence the function returns <code>6</code> as teh number of bits. The <a href="https://pkg.go.dev/math/bits@go1.18.1#OnesCount">OnesCount</a> Function is also similar but it returns the number of set bits(the <code>1</code> bit) in the number provided to it.</p>
<p>We'll see this sub package in a separate section of its own. Bits is really a great pacakge to work with bits and low level manipulation of numbers in Golang.</p>
<h2>The Complex package</h2>
<p>The complex subpackage is really specific to the operation to the complex numbers and its operations. Using complex numbers with basic operations and trigonometric functions are provided in the package.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;math/cmplx&quot;
)

func main() {

	x := complex(5, 8)
	y := complex(3, 4)
	mod_x := cmplx.Abs(x)
	mod_y := cmplx.Abs(y)
	conj_x := cmplx.Conj(x)
	phase_x := cmplx.Phase(x)
	mod, phase := cmplx.Polar(x)

	fmt.Println(&quot;x = &quot;, x)
	fmt.Println(&quot;Modulus of x = &quot;, mod_x)
	fmt.Println(&quot;Modulus of y = &quot;, mod_y)
	fmt.Println(&quot;Conjugate of x = &quot;, conj_x)
	fmt.Println(&quot;Phase of x = &quot;, phase_x)
	fmt.Printf(&quot;Polar Form : %v, %v
&quot;, mod, phase)

}
</code></pre>
<pre><code>$ go run basic-functions/complex.go
x =  (5+8i)
Modulus of x =  9.433981132056603
Modulus of y =  5
Conjugate of x =  (5-8i)
Phase of x =  1.0121970114513341
Polar Form : 9.433981132056603, 1.0121970114513341

</code></pre>
<p>We have used the complex function to create complex numbers. The <code>cmplx</code> subpackage in the math package provides many functions to play with trignometric and simple operations with complex numbers. The <a href="https://pkg.go.dev/math/cmplx@go1.18.1#Abs">Abs</a> function is used to get the modulus of the provided complex number. The modulus is calculated with <code>sqrt(x^2 + y^2)</code>, this gives the magnitude of the complex number. Here, we get the modulus as <code>9.43</code> as <code>sqrt(25 + 64)</code> for the complex number <code>5+8i</code>. Also, for <code>3+4i</code> the modulus becomes <code>sqrt(9+16)</code> which turns out to be <code>5</code>. The <a href="https://pkg.go.dev/math/cmplx@go1.18.1#Conj">Conjugate</a> function is used to get the conjugate of the provided complex number.</p>
<p>Also the phase or the Argument of the complex number can be obtained with the <a href="https://pkg.go.dev/math/cmplx@go1.18.1#Phase">Phase</a> function. The phase is caluculated by the formula <code>tan^-1 (y/x)</code> but the angle is returned in randians. So for <code>x = 5+8i</code> the argument/Phase becomes <code>tan^-1( 8/5)</code> which is <code>57.995 degrees</code> or <code>1.012 radians</code>.</p>
<p>We have the <a href="https://pkg.go.dev/math/cmplx@go1.18.1#Polar">Polar</a> function which gives the polar form of the complex number i.e. <code>(modulus r, phase theta)</code> So this function returns two values the modulus and the argument/phase of the complex number. We have already calcualted both the values but this functions gets both of them in a single function. Quite neat, we can even ignore one value after the return of the function by using the ignore operator <code>_, phase := cmplx.Polar(5+7i)</code> to only care and get the phase/argument of the complex number or <code>modulus, _ := cmplx.Polar(5+7i)</code> to get the modulus from the complex number.</p>
<p>So that's some basic operations on complex numbers, this might have very few use cases but it's still quite useful when needed.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So from this section we were able to get a bit deeper introduction to the <code>math</code> package in golang's standard library. We covered some few important functions and constants in the main math package along with the glimpse of other subpackages like <code>rand</code>, <code>cmplx</code> and <code>bits</code>. We didn't get too much in detail with those sub packages as they can be explored on a separate section of their own. Hopefully, you have got a godd overview of the math package in golang which again is really important aspect in programming.</p>
<p>Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
