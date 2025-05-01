---
article_html: "<h2>Introduction</h2>\n<p>One of the key features that set Go apart
  from many other languages is its native support for Goroutines - lightweight concurrent
  functions that can run concurrently and efficiently manage concurrency tasks. Goroutines
  are an essential aspect of Go's concurrency model, enabling developers to build
  highly concurrent and performant applications effortlessly.</p>\n<p>In this blog
  post, we will dive into the world of Goroutines and explore how they work, why they
  are essential for concurrent programming in Go, and how they can greatly improve
  the responsiveness and efficiency of your applications. This post will cover go
  routines and a primer on wait groups, in the next article we will be looking deeply
  into channels where all these three things can be demonstrated and understood in
  a better and more useful way.</p>\n<h2>Concurrency and Parallelism</h2>\n<p>This
  two concepts are quite crucial before diving into the fundamentals of concurreny.</p>\n<p><strong>Concurrency</strong>
  is about dealing with multiple things at once. Concurrent programs can have several
  processes or threads running simultaneously on a single CPU core by rapidly switching
  between them (context switching). The threads are interleaved, not necessarily executing
  at literally the same time. The CPU can switch between these tasks to give the appearance
  of simultaneous progress.</p>\n<p><strong>Parallelism</strong> is about doing multiple
  things at literally the same time. Parallel programs can execute multiple computations
  simultaneously on separate CPU cores. The threads actually execute in parallel.</p>\n<h2>What
  is a Go Routine</h2>\n<p>A go routine is a simple lightweight thread managed by
  the Go runtime. In the simplest of terms, a go routine can be defined as:</p>\n<blockquote>\n<p>Go
  routine is a way to perform multiple tasks within a program, allowing different
  parts of the program to work simultaneously and make the most out of the resources.</p>\n</blockquote>\n<p>Also
  it can be stated as:</p>\n<p>A goroutine in Golang is a lightweight, independently
  executing function that runs concurrently with other goroutines within the same
  address space. In other words, it is a concurrent unit of execution.</p>\n<p>Focussing
  on the word <strong>same address space</strong> that will be really critical in
  the later sections of this article.</p>\n<h2>Features of Go Routines</h2>\n<p>Go
  routines form a key part of Go's concurrency model. Here are some of the key features
  of go routines:</p>\n<ul>\n<li>\n<p><strong>Lightweight Thread</strong>:\nA Go routine
  is often referred to as a lightweight thread.</p>\n</li>\n<li>\n<p><strong>Independent
  Execution</strong>:\nGo routines run independently of each other, enabling concurrent
  execution.</p>\n</li>\n<li>\n<p><strong>Managed by Go</strong>:\nGo routines are
  managed by the Go runtime, making them easy to use.</p>\n</li>\n<li>\n<p><strong>Low
  Overhead</strong>:\nGo routines have low memory overhead, allowing us to create
  thousands of them efficiently.</p>\n</li>\n<li>\n<p><strong>Communication</strong>:\nGo
  routines can communicate and synchronize data through channels.</p>\n</li>\n<li>\n<p><strong>Asynchronous</strong>:\nGo
  routines can execute asynchronously, allowing other parts of the program to continue
  running.</p>\n</li>\n<li>\n<p><strong>Scalability</strong>:\nGo routines are the
  foundation of scalable concurrent programming in Go.</p>\n</li>\n</ul>\n<p>Unlike
  threads in other languages, goroutines are cheap and you can easily create thousands
  or even millions of them in a program.</p>\n<h2>Example of Go Routines</h2>\n<p>Creating
  a go routine is not hard, just add the keyword <code>go</code> before the function
  call and the go runtime will create a new go routine inside the main function or
  wherever is the context from. Remember, the main function is also a go routine.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;time&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">process</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Hello World!&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Sleep</span><span class=\"p\">(</span><span
  class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nx\">Second</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">start</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Now</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"nx\">process</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"nx\">process</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">end</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Now</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">duration</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">end</span><span
  class=\"p\">.</span><span class=\"nx\">Sub</span><span class=\"p\">(</span><span
  class=\"nx\">start</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">duration</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>scripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\nHello World!\n<span class=\"m\">15</span>.607µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\n<span class=\"m\">9</span>.889µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\n<span class=\"m\">8</span>.834µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\n<span class=\"m\">9</span>.158µs\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\n<span class=\"m\">12</span>.54µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\nHello World!\nHello World!\n<span
  class=\"m\">10</span>.19µs\n\n\nscripts/go-routines on  main via \U0001F439 v1.20
  \n$ go run main.go \n\n<span class=\"m\">14</span>.1µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ \n</pre></div>\n\n</pre>\n\n<p>Quite unpredictable
  output right? This is a power of go-routines. It is <strong>asynchronous</strong>
  so it will not block the main function. The two function calls to <code>process</code>,
  are executed independently of the main function scope. The program just calls the
  process function and captures the output, and sequentially reaches the end of the
  program(main function), At this point the go routines(threads) inside the main function
  are stopped abruptly.</p>\n<p>Let's break it down if it is not clear yet.</p>\n<ul>\n<li>The
  main function starts.</li>\n<li><code>go process()</code> creates a Go routine and
  starts its execution.</li>\n<li>Meanwhile, another call to <code>go process()</code>
  creates a separate Go routine and starts its execution as well.</li>\n<li>Meanwhile,
  calculate the time difference between the main function start and end.</li>\n<li>Meanwhile
  main function ends.</li>\n</ul>\n<p>So, in summary, the main function is only able
  to capture the output of the <code>duration</code> as it is synchronous and if the
  process has been executed, it prints the <code>Hello World!</code> message. Hence
  we have different outputs because the uncontrolled concurrency, lack of coordination,
  and OS scheduling across the program runs are different.</p>\n<h2>Wait Groups</h2>\n<p>In
  simple terms, WaitGroup is used to synchronize multiple goroutines and to wait for
  them to finish executing. This allows the go routines to be completed before the
  completion of the main function, hence it blocks the main function from leaving/exiting
  the scope.</p>\n<blockquote>\n<p>A WaitGroup is a synchronization primitive that
  allows a goroutine to wait for a collection of other goroutines to finish executing.</p>\n</blockquote>\n<ul>\n<li>A
  WaitGroup is initialized with a counter representing the number of goroutines to
  wait for.</li>\n<li>The Add() method increments the counter by the given value.
  This is called by each goroutine to indicate it is running.</li>\n<li>The main goroutine
  calls Add() to set the initial count, then launches worker goroutines.</li>\n<li>A
  WaitGroup is typically passed by a pointer to goroutines that need to be waited
  on.</li>\n<li>The Done() method decrements the counter by 1. Goroutines call this
  when finished.</li>\n<li>Each worker calls Done() when finished, decrementing the
  counter.</li>\n<li>The Wait() method blocks until the counter reaches 0, indicating
  all goroutines have finished.</li>\n<li>Main calls Wait() to block until Done()
  brings counter to 0.</li>\n</ul>\n<p>This provides a simple way to synchronize multiple
  goroutines finishing their work with a main thread that needs to wait for them to
  complete. The counter ensures the main thread knows how many goroutines it is waiting
  for. Interacting and working with go routines with synchronization using a wait
  group is quite intuitive and simple to follow. Let's look at a simple example below:</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;time&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">process</span><span class=\"p\">(</span><span class=\"nx\">pid</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Started process %d\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">pid</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nx\">Sleep</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"w\"> </span><span class=\"nx\">time</span><span class=\"p\">.</span><span
  class=\"nx\">Second</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;Completed
  process %d\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">pid</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">now</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">time</span><span class=\"p\">.</span><span
  class=\"nx\">Now</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">i</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">;</span><span class=\"w\">
  </span><span class=\"nx\">i</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">10</span><span class=\"p\">;</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">go</span><span class=\"w\"> </span><span
  class=\"nx\">process</span><span class=\"p\">(</span><span class=\"nx\">i</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">wg</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Wait</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;All
  processes completed&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">end</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Now</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">end</span><span class=\"p\">.</span><span class=\"nx\">Sub</span><span
  class=\"p\">(</span><span class=\"nx\">now</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>scripts/go-routines
  on  main <span class=\"o\">[</span>?<span class=\"o\">]</span> via \U0001F439 v1.20
  \n$ go run wg.go \nStarted process <span class=\"m\">9</span>\nStarted process <span
  class=\"m\">0</span>\nStarted process <span class=\"m\">1</span>\nStarted process
  <span class=\"m\">2</span>\nStarted process <span class=\"m\">3</span>\nStarted
  process <span class=\"m\">4</span>\nStarted process <span class=\"m\">5</span>\nStarted
  process <span class=\"m\">6</span>\nStarted process <span class=\"m\">7</span>\nStarted
  process <span class=\"m\">8</span>\nCompleted process <span class=\"m\">8</span>\nCompleted
  process <span class=\"m\">3</span>\nCompleted process <span class=\"m\">9</span>\nCompleted
  process <span class=\"m\">0</span>\nCompleted process <span class=\"m\">1</span>\nCompleted
  process <span class=\"m\">2</span>\nCompleted process <span class=\"m\">5</span>\nCompleted
  process <span class=\"m\">4</span>\nCompleted process <span class=\"m\">6</span>\nCompleted
  process <span class=\"m\">7</span>\n<span class=\"m\">1</span>.000563257s\nAll processes
  completed\n\nscripts/go-routines on  main <span class=\"o\">[</span>?<span class=\"o\">]</span>
  via \U0001F439 v1.20 \n$ \n</pre></div>\n\n</pre>\n\n<p>In the above example, we
  have used the same function <code>process</code> but with a slight twist, we have
  added a process id, just an integer to represent the go routine. We print the start
  and completion of the function in between the sleep for 1 second. We also have a
  wait group.</p>\n<ul>\n<li>A <a href=\"https://pkg.go.dev/sync#WaitGroup\">WaitGroup</a>
  is basically a struct type defined in the <a href=\"https://pkg.go.dev/sync\">sync</a>
  package.</li>\n<li>The variable <code>wg</code> is a new wait group instance that
  will be used to synchronize and wait for the completion of groups of go routines.</li>\n</ul>\n<p>We
  create a for loop with <code>10</code> iterations, so as to spawn 10 <code>process</code>
  function calls. We first use the <code>wg.Add(1)</code> which says to the wait group
  to wait for 1 go routine. The immediate next line is a go routine <code>go process()</code>
  which takes in a <code>pid</code> just to track which go routine is being executed
  in the loop.</p>\n<p>Inside the <code>process</code> function, we simply say the
  process with the given <code>pid</code> has started as a print statement, sleeps
  for a second, and then print the end of the process with <code>pid</code>. The end
  of the function is marked with <code>defer wg.Done()</code>, this is to indicate
  that the go routine is completed.</p>\n<p>The wg(wait group) has a counter that
  keeps track of the number of go routines it has to synchronize or wait for them
  to complete. In the <a href=\"https://pkg.go.dev/sync#WaitGroup.Add\">Add</a> function,
  the internal counter in the <code>WaitGroup</code> is incremented by the <code>delta</code>
  the integer parsed as a parameter. And the <a href=\"https://pkg.go.dev/sync#WaitGroup.Done\">Done</a>
  function, decrements the internal counter in the <code>WaitGroup</code> which indicates
  that the go routine is completed.</p>\n<p>In the main function, we call the <a href=\"https://pkg.go.dev/sync#WaitGroup.Wait\">wg.Wait</a>
  which will block until the counter for the <code>WaitGroup</code> is 0 i.e. all
  go routines have completed their execution. So, we created 10 go routines which
  ran concurrently but were synchronized with the help of WaitGroups. The WaitGroup
  is allowed to block the main function till all the go routines are done executing.</p>\n<h2>Go
  Routine with WaitGroup Use Cases</h2>\n<p>Go routines can be used for creating asynchronous
  tasks, and also for creating concurrent tasks. By using wait groups, we can create
  a way to wait for multiple goroutines to complete. By using go routines and wait
  groups, we can complete n number of tasks in a time frame of 1 task's completion.
  However, to create concurrent communication between the other processes, we need
  <code>channels</code> (which we will explore in the next article).</p>\n<p>Here
  is a simple breakdown of what asynchronous and concurrent tasks might refer to:</p>\n<p><strong>Asynchronous</strong>
  tasks run independently of the main program flow, allowing the main program to continue
  executing without waiting for the task's completion. For example, not blocking other
  tasks in the sequential flow of the main function.</p>\n<p><strong>Concurrent</strong>
  tasks run simultaneously and can execute at the same time as other tasks. They make
  use of multiple threads (goroutines in Go's case) to achieve parallel execution.
  For example, running multiple tasks parallel will cut the time down for spinning
  each task after the completion of another task.</p>\n<p>Some of the asynchronous
  tasks that can be done with go routines might include the following:</p>\n<ul>\n<li>Sending
  a mail while saving the user to the database.</li>\n<li>Fetch and process data from
  multiple websites(web scraping/crawling).</li>\n<li>High-performance message brokers
  and queuing systems for inter-process communication.</li>\n</ul>\n<p>One practical
  way of using go routines with wait groups would be to send mail, we are not going
  to see the actual implementation of the mail-sending stuff. However, we can try
  to mimic how the setup might be for sending bulk mail. By creating a wait group
  and a list of mail addresses to send, a go routine can be created with a function
  that sends those emails.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">sendMail</span><span class=\"p\">(</span><span class=\"nx\">address</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Sending mail to&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">address</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Done</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"c1\">// Actual mail sending, smtp stuff</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// handle errors</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"c1\">// client,
  err := smtp.Dial(&quot;smtp.example.com:587&quot;)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"c1\">// errr = client.Mail(&quot;sender@example.com&quot;)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// err = client.Rcpt(address)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// wc, err
  := client.Data()</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"c1\">//_, err = wc.Write([]byte(&quot;This is the email body.&quot;))</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// err = wc.Close()</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// client.Quit()</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">emails</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"s\">&quot;recipient1@example.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"s\">&quot;recipient2@example.com&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"s\">&quot;xyz@example.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Add</span><span class=\"p\">(</span><span class=\"nb\">len</span><span
  class=\"p\">(</span><span class=\"nx\">emails</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">email</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">emails</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"nx\">sendMail</span><span class=\"p\">(</span><span
  class=\"nx\">email</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">wg</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Wait</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;All emails queued for sending&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// Do other stuff</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the <code>emails</code> is a list of email ids to send the mail.
  We have created a wait group that we can initialize with the total number of go
  routines probably to be executed. The <code>wg.Add</code> method is parsed with
  the number of emails to be sent hence that equals the go routines to spawn.</p>\n<p>So,
  we can then in the for loop, iterate over each mail and send the emails with the
  <code>sendMail</code> function as go routine. The <code>wg.Wait</code> function
  outside the loop, will make sure the main function is halted before all the go routines
  complete their execution.</p>\n<p>There is one more way to call a function as a
  go routine without changing its signature since we had to pass <code>wg</code> pointer
  as the WaitGroup reference to it to acknowledge the completion of the go routine.
  We could wrap these two operations wiz. calling the function and calling the <code>wg.Done</code>
  method with an anonymous function.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">sendMail</span><span class=\"p\">(</span><span class=\"nx\">address</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Sending mail to&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">address</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">emails</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"s\">&quot;recipient1@example.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"s\">&quot;recipient2@example.com&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"s\">&quot;xyz@example.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Add</span><span class=\"p\">(</span><span class=\"nb\">len</span><span
  class=\"p\">(</span><span class=\"nx\">emails</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">email</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">emails</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">mail</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">email</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"kd\">func</span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t\t</span><span class=\"nx\">sendMail</span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Done</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"p\">}(</span><span class=\"nx\">mail</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Wait</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;All emails queued for sending&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// Do other stuff</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This does the exact same thing,
  but we don't have to change the signature of the function, This keeps the functional
  logic of the function in its place and handles the concurrency on the go.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">email</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">emails</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">mail</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">email</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">go</span><span class=\"w\"> </span><span
  class=\"kd\">func</span><span class=\"p\">(</span><span class=\"nx\">m</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">sendMail</span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Done</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}(</span><span class=\"nx\">mail</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>If
  the above bit scares you, don't worry it's too simple.</p>\n<ul>\n<li>We are iterating
  over the email slice using a for loop and creating a Go routine for each email address.
  The loop variable email represents the email address at the current iteration.</li>\n<li>However,
  to avoid the loop variable capture issue (where all Go routines would share the
  same email variable), we create a new variable mail and assign the value of email
  to it. This step ensures that each Go routine captures its own copy of the email
  address.</li>\n<li>We immediately create an anonymous function (a closure) using
  the go keyword. This anonymous function takes the <code>mail</code> variable as
  a parameter m and is executed concurrently as a Go routine. Inside the Go routine,
  we call the <code>sendMail</code> function with the email address <code>m</code>.</li>\n<li>After
  the <code>sendMail</code> call has been executed i.e. email has been sent, we call
  wg.Done() to notify the WaitGroup that the Go routine has completed its work. This
  allows the WaitGroup to properly synchronize and wait for all Go routines to finish
  before the program proceeds beyond wg.Wait() in the main function.</li>\n</ul>\n<p>This
  is one way to do it if you want to separate the mail-sending logic from the goroutines/concurrency
  task. However, this should be handled with care as the variables inside the closure
  might be shared among all the goroutines instead of having individual variable literals.</p>\n<p>To
  ensure that each goroutine operates on its own copy of the email address, we use
  the approach of creating a new variable mail and passing it as a parameter to the
  anonymous function. This way, each goroutine captures its unique email address,
  avoiding any interference or unintended sharing of data between goroutines.</p>\n<h2>Mutual
  Exclusion Locks</h2>\n<p>In the previous examples, we saw how goroutines and wait
  groups allow us to run multiple tasks concurrently in Go. However, sometimes these
  concurrent goroutines need to access shared resources like memory, files, network
  sockets, etc.</p>\n<p>When more than one goroutine tries to access a resource at
  the same time, it can lead to <strong>race conditions</strong> and unpredictable
  behavior. To handle this, we need a way to ensure only one goroutine can access
  the resource at a time.</p>\n<p>This is where mutual exclusion locks come in. A
  <strong>mutual exclusion lock</strong>, or <strong>mutex</strong>, provides a mechanism
  to lock access to a shared resource. It ensures only one goroutine at a time can
  acquire the lock, blocking other goroutines until the lock is released.</p>\n<p>For
  example, say we have multiple goroutines trying to append data to the same memory
  buffer(could be file/database/etc.) concurrently:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span class=\"nx\">filename</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">contents</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">buffer</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Done</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"o\">*</span><span
  class=\"nx\">buffer</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"o\">*</span><span class=\"nx\">buffer</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">contents</span><span class=\"p\">)</span><span
  class=\"o\">...</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"nx\">filename</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">buffer</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mo\">0644</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">sharedBuffer</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">byte</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"mi\">2</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">go</span><span class=\"w\"> </span><span
  class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span class=\"s\">&quot;data/f1.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello &quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">sharedBuffer</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">wg</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;data/f1.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;World! &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">sharedBuffer</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">wg</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Wait</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nb\">string</span><span
  class=\"p\">(</span><span class=\"nx\">sharedBuffer</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run --race no-mutex.go\n<span class=\"o\">==================</span>\nWARNING:
  DATA RACE\nRead at 0x00c000012030 by goroutine <span class=\"m\">8</span>:\n  main.WriteToFile<span
  class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:11
  +0xe7\n  main.main.func2<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:24
  +0x64\n\nPrevious write at 0x00c000012030 by goroutine <span class=\"m\">7</span>:\n
  \ main.WriteToFile<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:11
  +0x16a\n  main.main.func1<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:23
  +0x64\n\nGoroutine <span class=\"m\">8</span> <span class=\"o\">(</span>running<span
  class=\"o\">)</span> created at:\n  main.main<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:24
  +0x1f6\n\nGoroutine <span class=\"m\">7</span> <span class=\"o\">(</span>finished<span
  class=\"o\">)</span> created at:\n  main.main<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:23
  +0x14e\n<span class=\"o\">==================</span>\n</pre></div>\n\n</pre>\n\n<p>This
  is because we are trying to access the same memory address at the same time. This
  is a race condition and can lead to unpredictable behavior. Try removing the <code>--race</code>
  flag while running, in this little stupid example, it might not be obvious, but
  in complex and constrained environments, this can get the application in serious
  trouble.</p>\n<p><strong>NOTE</strong>: We are using <code>go run --race no-mutex.go</code>
  to check if there are any race conditions in the program. This is the <a href=\"https://go.dev/doc/articles/race_detector\">race
  detector</a> flag in the run command.</p>\n<p>To avoid this race condition, we need
  to add the mutex locks provided in the <a href=\"https://pkg.go.dev/sync#Mutex\">sync.Mutex</a>
  type. There are methods like <a href=\"\">Lock</a>, <a href=\"\">Unlock</a>, and
  <a href=\"\">TryLock</a> which help in locking access of the resource to a single
  entity at a given time.</p>\n<p>When a goroutine calls <code>Lock()</code> on a
  mutex, it acquires the lock. If the mutex is already locked by another goroutine,
  the calling goroutine will be blocked (put to sleep) until the lock becomes available.
  Once the lock is acquired successfully, the goroutine can proceed with its critical
  section, which is the part of the code that should not be executed concurrently
  by multiple goroutines.</p>\n<p>When a goroutine calls <code>Unlock()</code> on
  a mutex, it releases the lock. This allows other waiting goroutines to acquire the
  lock and execute their critical sections. It's essential to ensure that <code>Unlock()</code>
  is called after the critical section to release the mutex and avoid deadlocks. The
  critical section/shared resource should not be accessed after the release of this
  lock.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span class=\"nx\">filename</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">contents</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">buffer</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">mutex</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">Mutex</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">contentBytes</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">contents</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Lock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"o\">*</span><span
  class=\"nx\">buffer</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"o\">*</span><span class=\"nx\">buffer</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">contentBytes</span><span class=\"o\">...</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">OpenFile</span><span class=\"p\">(</span><span
  class=\"nx\">filename</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">O_APPEND</span><span
  class=\"p\">|</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">O_WRONLY</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mo\">0644</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">Write</span><span class=\"p\">(</span><span
  class=\"nx\">contentBytes</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Unlock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">mut</span><span class=\"w\"> </span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">Mutex</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">sharedBuffer</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">byte</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"mi\">2</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">go</span><span class=\"w\"> </span><span
  class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span class=\"s\">&quot;data/f1.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello Gophers!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">sharedBuffer</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">wg</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">mut</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">WriteToFile</span><span
  class=\"p\">(</span><span class=\"s\">&quot;data/f1.txt&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Welcome to Goworld!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">sharedBuffer</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">wg</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">mut</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Wait</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">sharedBuffer</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run --race mutex.go\nWelcome to Goworld!\nHello Gophers!\n\n$ go run --race mutex.go\nHello
  Gophers!\nWelcome to Goworld!\n\n$ go run --race mutex.go\nHello Gophers!\nWelcome
  to Goworld!\n\n$ go run --race mutex.go\nWelcome to Goworld!\nHello Gophers!\n</pre></div>\n\n</pre>\n\n<p>The
  above example is a preventive measure for race conditions with a mutex Lock on the
  shared resource used in the go routines.</p>\n<p>Let's break down the code step
  by step. First, we initialize a few variables:</p>\n<ul>\n<li>\n<p><code>wg</code>
  as <code>sync.WaitGroup</code>:\nThe <code>wg</code> is a waitgroup that will be
  used for the synchronization of the go-routines by blocking the main function to
  exit</p>\n</li>\n<li>\n<p><code>mut</code> as <code>sync.Mutex</code>:\nThe <code>mut</code>
  is a structure that internally holds a few integer values indicating either a blocked
  or unblocked state. The <code>sync.Mutex</code> has two private fields wiz. <code>state</code>
  and <code>sema</code>, the state holds the mutex state either <code>0</code>(unlocked)
  or <code>1</code> as locked. The <code>sema</code> field is a <code>uint32</code>
  that is used for blocking and signaling, it acts as a semaphore to manage blocking
  and unblocking goroutines trying to acquire the mutex.\nThis will be used to acquire
  <code>Lock</code> and <code>Unlock</code> on the shared resource while writing the
  data to the file or appending the data to the resource.</p>\n</li>\n<li>\n<p><code>sharedBuffer</code>
  as <code>[]byte</code>:\nThe <code>sharedBuffer</code> is the actually shared resource
  that will be used to hold the strings for keeping track of the data written to the
  file. It will be the resource that would require to lock for mutating its value(appending
  to the slice) among the go routines.</p>\n</li>\n</ul>\n<p>We add <code>2</code>
  to the <code>wg</code> to indicate to wait for the two go-routines to complete,
  in the next line we call two go routines as the function <code>WriteToFile</code>.
  The <code>WriteToFile</code> is a function that takes in quite a few parameters
  namely, the filename, the content to write, the reference to the sharedBuffer, waitgroup,
  and the mutex.</p>\n<p>Inside the function <code>WriteToFile</code>:</p>\n<ul>\n<li>We
  first <code>defer</code> the <code>waitgroups</code> as <code>Done</code> i.e. to
  call the <code>wg.Done</code> method at the end of the function call.</li>\n<li>Typecast
  the <code>contents</code> from <code>string</code> as the <code>[]byte</code>.</li>\n<li>Acquire
  the <code>mutex.Lock()</code> i.e. to say &quot;The below operations should not
  be done concurrently, one at a time&quot;. We then append the <code>contents</code>
  to the <code>buffer</code> which is a pointer to the <code>sharedBuffer</code> in
  the main function, so essentially we are trying to mutate the <code>sharedBuffer</code>
  in this function.</li>\n<li>Open the file as <code>O_APPEND</code> and <code>O_WRONLY</code>
  to indicate that the file should be opened in append/writing mode. (We have observed
  this type of operation in the <a href=\"https://www.meetgor.com/golang-file-write/\">Golang:
  File Write</a> article)</li>\n<li>Use the <code>Write</code> method to write the
  slice of bytes(contents) into the file that we opened. We have a <code>defer</code>
  for the closing of the file.</li>\n</ul>\n<p>We obviously check for errors and print
  them, but it could be a <code>panic</code> or <code>log</code> depending on the
  type of operation the application is doing.\nSo that is the all operation we want
  to do, so we finally open the lock with <code>mutex.Unlock()</code> which will allow
  the other go routine if any to access the resource and proceed with the operations.</p>\n<h3>Read
  Write Mutual Exclusion Lock</h3>\n<p>The mutual exclusion lock is good if you have
  a write operation-heavy application. However, if we have read write operations in
  more or the less in same proportion(read heavy) we don't want to have readers getting
  blocked when other readers are accessing the resource since it is not a mutation
  process.</p>\n<p>We could allow many readers to read simultaneously. But for the
  writing operation, we want to block the readers/writers. The writer should be given
  preference first in case of a lock by other writers. This would prevent a writer
  from waiting for readers to complete. This is usually referred to as <strong>Read
  Write Mutual Exclusion</strong>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;time&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">reader</span><span class=\"p\">(</span><span class=\"nx\">id</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">mutex</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">Mutex</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"p\">!</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">TryLock</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Reader %d blocked!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Unlock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Reader %d: read count %d\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">id</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">writer</span><span
  class=\"p\">(</span><span class=\"nx\">id</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">increment</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">count</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"kt\">int</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">mutex</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">Mutex</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Done</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">mutex</span><span class=\"p\">.</span><span
  class=\"nx\">Unlock</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">mutex</span><span class=\"p\">.</span><span
  class=\"nx\">Lock</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"w\"> </span><span class=\"o\">+=</span><span class=\"w\"> </span><span class=\"nx\">increment</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Sleep</span><span class=\"p\">(</span><span
  class=\"mi\">5</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"w\">
  </span><span class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nx\">Millisecond</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Writer %d: wrote count %d\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">1</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">mutex</span><span class=\"w\"> </span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">Mutex</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">readers</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">5</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">writers</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">3</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Add</span><span class=\"p\">(</span><span
  class=\"nx\">readers</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">i</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">;</span><span class=\"w\">
  </span><span class=\"nx\">i</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">readers</span><span class=\"p\">;</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"o\">&amp;</span><span class=\"nx\">count</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">wg</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">mutex</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Add</span><span class=\"p\">(</span><span class=\"nx\">writers</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">0</span><span
  class=\"p\">;</span><span class=\"w\"> </span><span class=\"nx\">i</span><span class=\"w\">
  </span><span class=\"p\">&lt;</span><span class=\"w\"> </span><span class=\"nx\">writers</span><span
  class=\"p\">;</span><span class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">writer</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">count</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">wg</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">mutex</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Wait</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run --race no-rwmutex.go \nReader <span class=\"m\">0</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">3</span> blocked!\nReader
  <span class=\"m\">2</span> blocked!\nReader <span class=\"m\">1</span> blocked!\nReader
  <span class=\"m\">4</span> blocked!\nWriter <span class=\"m\">0</span>: wrote count
  <span class=\"m\">2</span>\nWriter <span class=\"m\">1</span>: wrote count <span
  class=\"m\">3</span>\nWriter <span class=\"m\">2</span>: wrote count <span class=\"m\">4</span>\n\n$
  go run --race no-rwmutex.go \nReader <span class=\"m\">2</span> blocked!\nReader
  <span class=\"m\">0</span>: <span class=\"nb\">read</span> count <span class=\"m\">1</span>\nReader
  <span class=\"m\">1</span> blocked!\nReader <span class=\"m\">4</span> blocked!\nReader
  <span class=\"m\">3</span> blocked!\nWriter <span class=\"m\">2</span>: wrote count
  <span class=\"m\">2</span>\nWriter <span class=\"m\">0</span>: wrote count <span
  class=\"m\">3</span>\nWriter <span class=\"m\">1</span>: wrote count <span class=\"m\">4</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above example has a <code>reader</code> and a <code>writer</code> method, the <code>reader</code>
  method simply has to read the shared resource <code>count</code>. It acquires a
  <code>Mutex</code> lock before reading and unlocks it thereafter. Similarly, the
  <code>writer</code> function is used for incrementing the <code>count</code> shared
  resource.</p>\n<p>The <code>reader</code> method has a <a href=\"https://pkg.go.dev/sync#Mutex.TryLock\">TryLock</a>
  method that tries to acquire a mutex lock on the resource, if the resource is already
  locked, the function will return <code>false</code> and hence we can say that the
  reading is blocked(just for demonstration). And if the function <code>TryLock</code>
  returns <code>true</code>, it will acquire the <code>Lock</code>. We further <code>defer</code>
  the <code>Unlock</code> and access the <code>count</code> variable which is passed
  as a reference.</p>\n<p>The <code>writer</code> method is simply acquiring the <code>Lock</code>
  and incrementing the <code>counter</code> and thereafter <code>Unlock</code> is
  called with <code>defer</code>.</p>\n<p>In the above code:</p>\n<ul>\n<li>The <code>reader</code>
  and the <code>writer</code> both might be waiting for the lock to be released, however,
  for readers to wait for reading doesn't make sense.</li>\n<li>Because if you would
  want to just read a particular memory address, there shouldn't be any locks for
  one reader to wait for other readers to finish.</li>\n<li>However, for writing,
  there has to be a lock. The mutex lock will lock the resource irrespective of the
  <code>reader</code> or <code>writer</code>.</li>\n</ul>\n<p>This might not be as
  visible here, but it might be the reason, that all the <code>readers</code> are
  blocked from reading due to another reader's or writer's lock.</p>\n<p>The <code>sync</code>
  package has the <a href=\"https://pkg.go.dev/sync#RWMutex\">RWMutex</a> that does
  this exact same thing. It is almost similar to the <code>Mutex</code> however, it
  would allow concurrent reading operation and prefer writing operation before readers
  to prevent writer starvation.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;time&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">reader</span><span class=\"p\">(</span><span class=\"nx\">id</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">mutex</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">RWMutex</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">RUnlock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"p\">!</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">TryRLock</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Reader %d blocked!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Reader %d: read count %d\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">writer</span><span class=\"p\">(</span><span class=\"nx\">id</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">increment</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">mutex</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">RWMutex</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Unlock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Lock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"o\">*</span><span
  class=\"nx\">count</span><span class=\"w\"> </span><span class=\"o\">+=</span><span
  class=\"w\"> </span><span class=\"nx\">increment</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">time</span><span class=\"p\">.</span><span
  class=\"nx\">Sleep</span><span class=\"p\">(</span><span class=\"mi\">5</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"w\"> </span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Millisecond</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Writer %d: wrote count %d\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">id</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">1</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">mutex</span><span class=\"w\"> </span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">RWMutex</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">readers</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">5</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">writers</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">3</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Add</span><span class=\"p\">(</span><span
  class=\"nx\">readers</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">i</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">;</span><span class=\"w\">
  </span><span class=\"nx\">i</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">readers</span><span class=\"p\">;</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"o\">&amp;</span><span class=\"nx\">count</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">wg</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">mutex</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Add</span><span class=\"p\">(</span><span class=\"nx\">writers</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">0</span><span
  class=\"p\">;</span><span class=\"w\"> </span><span class=\"nx\">i</span><span class=\"w\">
  </span><span class=\"p\">&lt;</span><span class=\"w\"> </span><span class=\"nx\">writers</span><span
  class=\"p\">;</span><span class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">writer</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">count</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">wg</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">mutex</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Wait</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run --race rwmutex.go \nReader <span class=\"m\">0</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">3</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">1</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">2</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">4</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nWriter <span class=\"m\">1</span>: wrote count
  <span class=\"m\">2</span>\nWriter <span class=\"m\">0</span>: wrote count <span
  class=\"m\">3</span>\nWriter <span class=\"m\">2</span>: wrote count <span class=\"m\">4</span>\nWriter
  <span class=\"m\">3</span>: wrote count <span class=\"m\">5</span>\nWriter <span
  class=\"m\">4</span>: wrote count <span class=\"m\">6</span>\n\n$ go run --race
  rwmutex.go \nReader <span class=\"m\">1</span> blocked!\nReader <span class=\"m\">0</span>:
  <span class=\"nb\">read</span> count <span class=\"m\">1</span>\nReader <span class=\"m\">2</span>
  blocked!\nReader <span class=\"m\">3</span> blocked!\nReader <span class=\"m\">4</span>
  blocked!\nWriter <span class=\"m\">4</span>: wrote count <span class=\"m\">2</span>\nWriter
  <span class=\"m\">0</span>: wrote count <span class=\"m\">3</span>\nWriter <span
  class=\"m\">1</span>: wrote count <span class=\"m\">4</span>\nWriter <span class=\"m\">3</span>:
  wrote count <span class=\"m\">5</span>\nWriter <span class=\"m\">2</span>: wrote
  count <span class=\"m\">6</span>\n</pre></div>\n\n</pre>\n\n<p>In the modified example,
  all the logic remains the same, just the <code>sync.Mutex</code> is replaced with
  <code>sync.RWMutex</code>. Also for trying to acquire the lock in the <code>reader</code>
  method <a href=\"https://pkg.go.dev/sync#RWMutex.RLock\">TryRLock</a> is used instead
  of <a href=\"https://pkg.go.dev/sync#RWMutex.TryRLock\">TryLock</a> which will check
  if the existing lock acquired is of a reader or writer, if it is a reader, it will
  return <code>true</code>, else <code>false</code>. Also the <code>Unlock</code>
  is replaced with the <code>RUnlock</code> method for releasing the read lock. In
  the <code>writer</code> method, everything remains the same so the writer has to
  acquire the lock irrespective of whether the current lock is from the reader/writer,
  so it is a normal <code>Lock</code> and <code>Unlock</code>.</p>\n<p>In the above
  example, we can see all the read operations sometimes are executed instead of getting
  blocked. This is due to the <code>RWMutex</code> Lock and Unlock on the read operation/function.</p>\n<ul>\n<li>When
  one reader is reading, it can't block other readers.</li>\n<li>However with simple
  <code>Mutex</code>, the reader is even blocked when another reader is reading.</li>\n<li>For
  the write operation, it will be blocked as usual, so if a writer is performing a
  write operation and a reader/readers come in, they will be blocked, also if in the
  meantime, while the resource is locked, another writer comes in, the writer will
  be given preference instead of waiting for all readers to complete. This prevents
  writer starvation.</li>\n</ul>\n<p>You can see, that the readers are still blocked
  the majority of the time, however, they are blocked due to the writer locking the
  resource and not any other reader.</p>\n<p>There is a bit more suitable <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/go-routines/file-rw.go\">example</a>
  related to <code>file</code> reading and writing with <code>RWMutex</code>, make
  sure to check that out to get a more clear understanding of an actual case of using
  the <code>RWMutex</code>.</p>\n<h3>Channels</h3>\n<p>This is a big part and I would
  like to delve into this topic in a separate post. There are some patterns like <code>fan-in</code>,
  <code>fan-out</code>, <code>worker-pool</code>, <code>pub-sub</code>, etc. which
  are really common in web applications and backend systems. These patterns we shall
  explore in the next article.</p>\n<p>Channels are a way to provide a safe and idiomatic
  way for Goroutines to exchange data and coordinate their execution without resorting
  to low-level mechanisms like shared memory or explicit locking.</p>\n<p>That's it
  from the 30th part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/go-routines\">100
  days of Golang</a> repository.</p>\n<h3>References:</h3>\n<ul>\n<li><a href=\"https://blog.nindalf.com/posts/how-goroutines-work/\">How
  Goroutines Work</a></li>\n<li><a href=\"https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/\">Concurrency
  patterns in Golang: WaitGroups and Goroutines</a></li>\n<li><a href=\"https://riteeksrivastava.medium.com/a-complete-journey-with-goroutines-8472630c7f5c\">A
  complete journey with Goroutines</a></li>\n</ul>\n<h2>Conclusion</h2>\n<p>From this
  part of the series, the fundamentals of golang's concurrency model were understood
  specifically spawning go-routines, synchronously executing go-routines with the
  help of a wait group, mutex locks, and how to secure concurrent access to a shared
  resource. In the next part of the series, these concepts will be used in asynchronous
  communication using channels.</p>\n<p>Hopefully, you have got the basics of concurrency
  in golang cleared from this post. If you have any queries, suggestions, or feedback,
  please feel free to comment below or contact me on the social handles. Thank you
  for reading. Happy Coding :)</p>\n"
cover: ''
date: 2023-07-29
datetime: 2023-07-29 00:00:00+00:00
description: Exploring Goroutines and WaitGroups in Go, understanding the fundamentals
  of Go's concurrency model.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2023-07-29-GO-30-GoRoutines.md
html: "<h2>Introduction</h2>\n<p>One of the key features that set Go apart from many
  other languages is its native support for Goroutines - lightweight concurrent functions
  that can run concurrently and efficiently manage concurrency tasks. Goroutines are
  an essential aspect of Go's concurrency model, enabling developers to build highly
  concurrent and performant applications effortlessly.</p>\n<p>In this blog post,
  we will dive into the world of Goroutines and explore how they work, why they are
  essential for concurrent programming in Go, and how they can greatly improve the
  responsiveness and efficiency of your applications. This post will cover go routines
  and a primer on wait groups, in the next article we will be looking deeply into
  channels where all these three things can be demonstrated and understood in a better
  and more useful way.</p>\n<h2>Concurrency and Parallelism</h2>\n<p>This two concepts
  are quite crucial before diving into the fundamentals of concurreny.</p>\n<p><strong>Concurrency</strong>
  is about dealing with multiple things at once. Concurrent programs can have several
  processes or threads running simultaneously on a single CPU core by rapidly switching
  between them (context switching). The threads are interleaved, not necessarily executing
  at literally the same time. The CPU can switch between these tasks to give the appearance
  of simultaneous progress.</p>\n<p><strong>Parallelism</strong> is about doing multiple
  things at literally the same time. Parallel programs can execute multiple computations
  simultaneously on separate CPU cores. The threads actually execute in parallel.</p>\n<h2>What
  is a Go Routine</h2>\n<p>A go routine is a simple lightweight thread managed by
  the Go runtime. In the simplest of terms, a go routine can be defined as:</p>\n<blockquote>\n<p>Go
  routine is a way to perform multiple tasks within a program, allowing different
  parts of the program to work simultaneously and make the most out of the resources.</p>\n</blockquote>\n<p>Also
  it can be stated as:</p>\n<p>A goroutine in Golang is a lightweight, independently
  executing function that runs concurrently with other goroutines within the same
  address space. In other words, it is a concurrent unit of execution.</p>\n<p>Focussing
  on the word <strong>same address space</strong> that will be really critical in
  the later sections of this article.</p>\n<h2>Features of Go Routines</h2>\n<p>Go
  routines form a key part of Go's concurrency model. Here are some of the key features
  of go routines:</p>\n<ul>\n<li>\n<p><strong>Lightweight Thread</strong>:\nA Go routine
  is often referred to as a lightweight thread.</p>\n</li>\n<li>\n<p><strong>Independent
  Execution</strong>:\nGo routines run independently of each other, enabling concurrent
  execution.</p>\n</li>\n<li>\n<p><strong>Managed by Go</strong>:\nGo routines are
  managed by the Go runtime, making them easy to use.</p>\n</li>\n<li>\n<p><strong>Low
  Overhead</strong>:\nGo routines have low memory overhead, allowing us to create
  thousands of them efficiently.</p>\n</li>\n<li>\n<p><strong>Communication</strong>:\nGo
  routines can communicate and synchronize data through channels.</p>\n</li>\n<li>\n<p><strong>Asynchronous</strong>:\nGo
  routines can execute asynchronously, allowing other parts of the program to continue
  running.</p>\n</li>\n<li>\n<p><strong>Scalability</strong>:\nGo routines are the
  foundation of scalable concurrent programming in Go.</p>\n</li>\n</ul>\n<p>Unlike
  threads in other languages, goroutines are cheap and you can easily create thousands
  or even millions of them in a program.</p>\n<h2>Example of Go Routines</h2>\n<p>Creating
  a go routine is not hard, just add the keyword <code>go</code> before the function
  call and the go runtime will create a new go routine inside the main function or
  wherever is the context from. Remember, the main function is also a go routine.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;time&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">process</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Hello World!&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Sleep</span><span class=\"p\">(</span><span
  class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nx\">Second</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">start</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Now</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"nx\">process</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"nx\">process</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">end</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Now</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">duration</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">end</span><span
  class=\"p\">.</span><span class=\"nx\">Sub</span><span class=\"p\">(</span><span
  class=\"nx\">start</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">duration</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>scripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\nHello World!\n<span class=\"m\">15</span>.607µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\n<span class=\"m\">9</span>.889µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\n<span class=\"m\">8</span>.834µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\n<span class=\"m\">9</span>.158µs\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\n<span class=\"m\">12</span>.54µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ go run main.go \n\nHello World!\nHello World!\n<span
  class=\"m\">10</span>.19µs\n\n\nscripts/go-routines on  main via \U0001F439 v1.20
  \n$ go run main.go \n\n<span class=\"m\">14</span>.1µs\n\n\nscripts/go-routines
  on  main via \U0001F439 v1.20 \n$ \n</pre></div>\n\n</pre>\n\n<p>Quite unpredictable
  output right? This is a power of go-routines. It is <strong>asynchronous</strong>
  so it will not block the main function. The two function calls to <code>process</code>,
  are executed independently of the main function scope. The program just calls the
  process function and captures the output, and sequentially reaches the end of the
  program(main function), At this point the go routines(threads) inside the main function
  are stopped abruptly.</p>\n<p>Let's break it down if it is not clear yet.</p>\n<ul>\n<li>The
  main function starts.</li>\n<li><code>go process()</code> creates a Go routine and
  starts its execution.</li>\n<li>Meanwhile, another call to <code>go process()</code>
  creates a separate Go routine and starts its execution as well.</li>\n<li>Meanwhile,
  calculate the time difference between the main function start and end.</li>\n<li>Meanwhile
  main function ends.</li>\n</ul>\n<p>So, in summary, the main function is only able
  to capture the output of the <code>duration</code> as it is synchronous and if the
  process has been executed, it prints the <code>Hello World!</code> message. Hence
  we have different outputs because the uncontrolled concurrency, lack of coordination,
  and OS scheduling across the program runs are different.</p>\n<h2>Wait Groups</h2>\n<p>In
  simple terms, WaitGroup is used to synchronize multiple goroutines and to wait for
  them to finish executing. This allows the go routines to be completed before the
  completion of the main function, hence it blocks the main function from leaving/exiting
  the scope.</p>\n<blockquote>\n<p>A WaitGroup is a synchronization primitive that
  allows a goroutine to wait for a collection of other goroutines to finish executing.</p>\n</blockquote>\n<ul>\n<li>A
  WaitGroup is initialized with a counter representing the number of goroutines to
  wait for.</li>\n<li>The Add() method increments the counter by the given value.
  This is called by each goroutine to indicate it is running.</li>\n<li>The main goroutine
  calls Add() to set the initial count, then launches worker goroutines.</li>\n<li>A
  WaitGroup is typically passed by a pointer to goroutines that need to be waited
  on.</li>\n<li>The Done() method decrements the counter by 1. Goroutines call this
  when finished.</li>\n<li>Each worker calls Done() when finished, decrementing the
  counter.</li>\n<li>The Wait() method blocks until the counter reaches 0, indicating
  all goroutines have finished.</li>\n<li>Main calls Wait() to block until Done()
  brings counter to 0.</li>\n</ul>\n<p>This provides a simple way to synchronize multiple
  goroutines finishing their work with a main thread that needs to wait for them to
  complete. The counter ensures the main thread knows how many goroutines it is waiting
  for. Interacting and working with go routines with synchronization using a wait
  group is quite intuitive and simple to follow. Let's look at a simple example below:</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;time&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">process</span><span class=\"p\">(</span><span class=\"nx\">pid</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Started process %d\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">pid</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nx\">Sleep</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"w\"> </span><span class=\"nx\">time</span><span class=\"p\">.</span><span
  class=\"nx\">Second</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;Completed
  process %d\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">pid</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">now</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">time</span><span class=\"p\">.</span><span
  class=\"nx\">Now</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">i</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">;</span><span class=\"w\">
  </span><span class=\"nx\">i</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">10</span><span class=\"p\">;</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">go</span><span class=\"w\"> </span><span
  class=\"nx\">process</span><span class=\"p\">(</span><span class=\"nx\">i</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">wg</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Wait</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;All
  processes completed&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">end</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Now</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">end</span><span class=\"p\">.</span><span class=\"nx\">Sub</span><span
  class=\"p\">(</span><span class=\"nx\">now</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>scripts/go-routines
  on  main <span class=\"o\">[</span>?<span class=\"o\">]</span> via \U0001F439 v1.20
  \n$ go run wg.go \nStarted process <span class=\"m\">9</span>\nStarted process <span
  class=\"m\">0</span>\nStarted process <span class=\"m\">1</span>\nStarted process
  <span class=\"m\">2</span>\nStarted process <span class=\"m\">3</span>\nStarted
  process <span class=\"m\">4</span>\nStarted process <span class=\"m\">5</span>\nStarted
  process <span class=\"m\">6</span>\nStarted process <span class=\"m\">7</span>\nStarted
  process <span class=\"m\">8</span>\nCompleted process <span class=\"m\">8</span>\nCompleted
  process <span class=\"m\">3</span>\nCompleted process <span class=\"m\">9</span>\nCompleted
  process <span class=\"m\">0</span>\nCompleted process <span class=\"m\">1</span>\nCompleted
  process <span class=\"m\">2</span>\nCompleted process <span class=\"m\">5</span>\nCompleted
  process <span class=\"m\">4</span>\nCompleted process <span class=\"m\">6</span>\nCompleted
  process <span class=\"m\">7</span>\n<span class=\"m\">1</span>.000563257s\nAll processes
  completed\n\nscripts/go-routines on  main <span class=\"o\">[</span>?<span class=\"o\">]</span>
  via \U0001F439 v1.20 \n$ \n</pre></div>\n\n</pre>\n\n<p>In the above example, we
  have used the same function <code>process</code> but with a slight twist, we have
  added a process id, just an integer to represent the go routine. We print the start
  and completion of the function in between the sleep for 1 second. We also have a
  wait group.</p>\n<ul>\n<li>A <a href=\"https://pkg.go.dev/sync#WaitGroup\">WaitGroup</a>
  is basically a struct type defined in the <a href=\"https://pkg.go.dev/sync\">sync</a>
  package.</li>\n<li>The variable <code>wg</code> is a new wait group instance that
  will be used to synchronize and wait for the completion of groups of go routines.</li>\n</ul>\n<p>We
  create a for loop with <code>10</code> iterations, so as to spawn 10 <code>process</code>
  function calls. We first use the <code>wg.Add(1)</code> which says to the wait group
  to wait for 1 go routine. The immediate next line is a go routine <code>go process()</code>
  which takes in a <code>pid</code> just to track which go routine is being executed
  in the loop.</p>\n<p>Inside the <code>process</code> function, we simply say the
  process with the given <code>pid</code> has started as a print statement, sleeps
  for a second, and then print the end of the process with <code>pid</code>. The end
  of the function is marked with <code>defer wg.Done()</code>, this is to indicate
  that the go routine is completed.</p>\n<p>The wg(wait group) has a counter that
  keeps track of the number of go routines it has to synchronize or wait for them
  to complete. In the <a href=\"https://pkg.go.dev/sync#WaitGroup.Add\">Add</a> function,
  the internal counter in the <code>WaitGroup</code> is incremented by the <code>delta</code>
  the integer parsed as a parameter. And the <a href=\"https://pkg.go.dev/sync#WaitGroup.Done\">Done</a>
  function, decrements the internal counter in the <code>WaitGroup</code> which indicates
  that the go routine is completed.</p>\n<p>In the main function, we call the <a href=\"https://pkg.go.dev/sync#WaitGroup.Wait\">wg.Wait</a>
  which will block until the counter for the <code>WaitGroup</code> is 0 i.e. all
  go routines have completed their execution. So, we created 10 go routines which
  ran concurrently but were synchronized with the help of WaitGroups. The WaitGroup
  is allowed to block the main function till all the go routines are done executing.</p>\n<h2>Go
  Routine with WaitGroup Use Cases</h2>\n<p>Go routines can be used for creating asynchronous
  tasks, and also for creating concurrent tasks. By using wait groups, we can create
  a way to wait for multiple goroutines to complete. By using go routines and wait
  groups, we can complete n number of tasks in a time frame of 1 task's completion.
  However, to create concurrent communication between the other processes, we need
  <code>channels</code> (which we will explore in the next article).</p>\n<p>Here
  is a simple breakdown of what asynchronous and concurrent tasks might refer to:</p>\n<p><strong>Asynchronous</strong>
  tasks run independently of the main program flow, allowing the main program to continue
  executing without waiting for the task's completion. For example, not blocking other
  tasks in the sequential flow of the main function.</p>\n<p><strong>Concurrent</strong>
  tasks run simultaneously and can execute at the same time as other tasks. They make
  use of multiple threads (goroutines in Go's case) to achieve parallel execution.
  For example, running multiple tasks parallel will cut the time down for spinning
  each task after the completion of another task.</p>\n<p>Some of the asynchronous
  tasks that can be done with go routines might include the following:</p>\n<ul>\n<li>Sending
  a mail while saving the user to the database.</li>\n<li>Fetch and process data from
  multiple websites(web scraping/crawling).</li>\n<li>High-performance message brokers
  and queuing systems for inter-process communication.</li>\n</ul>\n<p>One practical
  way of using go routines with wait groups would be to send mail, we are not going
  to see the actual implementation of the mail-sending stuff. However, we can try
  to mimic how the setup might be for sending bulk mail. By creating a wait group
  and a list of mail addresses to send, a go routine can be created with a function
  that sends those emails.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">sendMail</span><span class=\"p\">(</span><span class=\"nx\">address</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Sending mail to&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">address</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Done</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"c1\">// Actual mail sending, smtp stuff</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// handle errors</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"c1\">// client,
  err := smtp.Dial(&quot;smtp.example.com:587&quot;)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"c1\">// errr = client.Mail(&quot;sender@example.com&quot;)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// err = client.Rcpt(address)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// wc, err
  := client.Data()</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"c1\">//_, err = wc.Write([]byte(&quot;This is the email body.&quot;))</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// err = wc.Close()</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// client.Quit()</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">emails</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"s\">&quot;recipient1@example.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"s\">&quot;recipient2@example.com&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"s\">&quot;xyz@example.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Add</span><span class=\"p\">(</span><span class=\"nb\">len</span><span
  class=\"p\">(</span><span class=\"nx\">emails</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">email</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">emails</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"nx\">sendMail</span><span class=\"p\">(</span><span
  class=\"nx\">email</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">wg</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Wait</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;All emails queued for sending&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// Do other stuff</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the <code>emails</code> is a list of email ids to send the mail.
  We have created a wait group that we can initialize with the total number of go
  routines probably to be executed. The <code>wg.Add</code> method is parsed with
  the number of emails to be sent hence that equals the go routines to spawn.</p>\n<p>So,
  we can then in the for loop, iterate over each mail and send the emails with the
  <code>sendMail</code> function as go routine. The <code>wg.Wait</code> function
  outside the loop, will make sure the main function is halted before all the go routines
  complete their execution.</p>\n<p>There is one more way to call a function as a
  go routine without changing its signature since we had to pass <code>wg</code> pointer
  as the WaitGroup reference to it to acknowledge the completion of the go routine.
  We could wrap these two operations wiz. calling the function and calling the <code>wg.Done</code>
  method with an anonymous function.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">sendMail</span><span class=\"p\">(</span><span class=\"nx\">address</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Sending mail to&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">address</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">emails</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"s\">&quot;recipient1@example.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"s\">&quot;recipient2@example.com&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"s\">&quot;xyz@example.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Add</span><span class=\"p\">(</span><span class=\"nb\">len</span><span
  class=\"p\">(</span><span class=\"nx\">emails</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">email</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">emails</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">mail</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">email</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"kd\">func</span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t\t</span><span class=\"nx\">sendMail</span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Done</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"p\">}(</span><span class=\"nx\">mail</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Wait</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;All emails queued for sending&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// Do other stuff</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This does the exact same thing,
  but we don't have to change the signature of the function, This keeps the functional
  logic of the function in its place and handles the concurrency on the go.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">email</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">emails</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">mail</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">email</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">go</span><span class=\"w\"> </span><span
  class=\"kd\">func</span><span class=\"p\">(</span><span class=\"nx\">m</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">sendMail</span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Done</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}(</span><span class=\"nx\">mail</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>If
  the above bit scares you, don't worry it's too simple.</p>\n<ul>\n<li>We are iterating
  over the email slice using a for loop and creating a Go routine for each email address.
  The loop variable email represents the email address at the current iteration.</li>\n<li>However,
  to avoid the loop variable capture issue (where all Go routines would share the
  same email variable), we create a new variable mail and assign the value of email
  to it. This step ensures that each Go routine captures its own copy of the email
  address.</li>\n<li>We immediately create an anonymous function (a closure) using
  the go keyword. This anonymous function takes the <code>mail</code> variable as
  a parameter m and is executed concurrently as a Go routine. Inside the Go routine,
  we call the <code>sendMail</code> function with the email address <code>m</code>.</li>\n<li>After
  the <code>sendMail</code> call has been executed i.e. email has been sent, we call
  wg.Done() to notify the WaitGroup that the Go routine has completed its work. This
  allows the WaitGroup to properly synchronize and wait for all Go routines to finish
  before the program proceeds beyond wg.Wait() in the main function.</li>\n</ul>\n<p>This
  is one way to do it if you want to separate the mail-sending logic from the goroutines/concurrency
  task. However, this should be handled with care as the variables inside the closure
  might be shared among all the goroutines instead of having individual variable literals.</p>\n<p>To
  ensure that each goroutine operates on its own copy of the email address, we use
  the approach of creating a new variable mail and passing it as a parameter to the
  anonymous function. This way, each goroutine captures its unique email address,
  avoiding any interference or unintended sharing of data between goroutines.</p>\n<h2>Mutual
  Exclusion Locks</h2>\n<p>In the previous examples, we saw how goroutines and wait
  groups allow us to run multiple tasks concurrently in Go. However, sometimes these
  concurrent goroutines need to access shared resources like memory, files, network
  sockets, etc.</p>\n<p>When more than one goroutine tries to access a resource at
  the same time, it can lead to <strong>race conditions</strong> and unpredictable
  behavior. To handle this, we need a way to ensure only one goroutine can access
  the resource at a time.</p>\n<p>This is where mutual exclusion locks come in. A
  <strong>mutual exclusion lock</strong>, or <strong>mutex</strong>, provides a mechanism
  to lock access to a shared resource. It ensures only one goroutine at a time can
  acquire the lock, blocking other goroutines until the lock is released.</p>\n<p>For
  example, say we have multiple goroutines trying to append data to the same memory
  buffer(could be file/database/etc.) concurrently:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span class=\"nx\">filename</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">contents</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">buffer</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Done</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"o\">*</span><span
  class=\"nx\">buffer</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"o\">*</span><span class=\"nx\">buffer</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">contents</span><span class=\"p\">)</span><span
  class=\"o\">...</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"nx\">filename</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">buffer</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mo\">0644</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">sharedBuffer</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">byte</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"mi\">2</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">go</span><span class=\"w\"> </span><span
  class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span class=\"s\">&quot;data/f1.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello &quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">sharedBuffer</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">wg</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">go</span><span
  class=\"w\"> </span><span class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;data/f1.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;World! &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">sharedBuffer</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">wg</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Wait</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nb\">string</span><span
  class=\"p\">(</span><span class=\"nx\">sharedBuffer</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run --race no-mutex.go\n<span class=\"o\">==================</span>\nWARNING:
  DATA RACE\nRead at 0x00c000012030 by goroutine <span class=\"m\">8</span>:\n  main.WriteToFile<span
  class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:11
  +0xe7\n  main.main.func2<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:24
  +0x64\n\nPrevious write at 0x00c000012030 by goroutine <span class=\"m\">7</span>:\n
  \ main.WriteToFile<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:11
  +0x16a\n  main.main.func1<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:23
  +0x64\n\nGoroutine <span class=\"m\">8</span> <span class=\"o\">(</span>running<span
  class=\"o\">)</span> created at:\n  main.main<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:24
  +0x1f6\n\nGoroutine <span class=\"m\">7</span> <span class=\"o\">(</span>finished<span
  class=\"o\">)</span> created at:\n  main.main<span class=\"o\">()</span>\n      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:23
  +0x14e\n<span class=\"o\">==================</span>\n</pre></div>\n\n</pre>\n\n<p>This
  is because we are trying to access the same memory address at the same time. This
  is a race condition and can lead to unpredictable behavior. Try removing the <code>--race</code>
  flag while running, in this little stupid example, it might not be obvious, but
  in complex and constrained environments, this can get the application in serious
  trouble.</p>\n<p><strong>NOTE</strong>: We are using <code>go run --race no-mutex.go</code>
  to check if there are any race conditions in the program. This is the <a href=\"https://go.dev/doc/articles/race_detector\">race
  detector</a> flag in the run command.</p>\n<p>To avoid this race condition, we need
  to add the mutex locks provided in the <a href=\"https://pkg.go.dev/sync#Mutex\">sync.Mutex</a>
  type. There are methods like <a href=\"\">Lock</a>, <a href=\"\">Unlock</a>, and
  <a href=\"\">TryLock</a> which help in locking access of the resource to a single
  entity at a given time.</p>\n<p>When a goroutine calls <code>Lock()</code> on a
  mutex, it acquires the lock. If the mutex is already locked by another goroutine,
  the calling goroutine will be blocked (put to sleep) until the lock becomes available.
  Once the lock is acquired successfully, the goroutine can proceed with its critical
  section, which is the part of the code that should not be executed concurrently
  by multiple goroutines.</p>\n<p>When a goroutine calls <code>Unlock()</code> on
  a mutex, it releases the lock. This allows other waiting goroutines to acquire the
  lock and execute their critical sections. It's essential to ensure that <code>Unlock()</code>
  is called after the critical section to release the mutex and avoid deadlocks. The
  critical section/shared resource should not be accessed after the release of this
  lock.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span class=\"nx\">filename</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">contents</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">buffer</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">mutex</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">Mutex</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">contentBytes</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">contents</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Lock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"o\">*</span><span
  class=\"nx\">buffer</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"o\">*</span><span class=\"nx\">buffer</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">contentBytes</span><span class=\"o\">...</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">OpenFile</span><span class=\"p\">(</span><span
  class=\"nx\">filename</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">O_APPEND</span><span
  class=\"p\">|</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">O_WRONLY</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mo\">0644</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">Write</span><span class=\"p\">(</span><span
  class=\"nx\">contentBytes</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Unlock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">mut</span><span class=\"w\"> </span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">Mutex</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">sharedBuffer</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">byte</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"mi\">2</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">go</span><span class=\"w\"> </span><span
  class=\"nx\">WriteToFile</span><span class=\"p\">(</span><span class=\"s\">&quot;data/f1.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello Gophers!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">sharedBuffer</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">wg</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">mut</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">WriteToFile</span><span
  class=\"p\">(</span><span class=\"s\">&quot;data/f1.txt&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Welcome to Goworld!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">sharedBuffer</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">wg</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">mut</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Wait</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">sharedBuffer</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run --race mutex.go\nWelcome to Goworld!\nHello Gophers!\n\n$ go run --race mutex.go\nHello
  Gophers!\nWelcome to Goworld!\n\n$ go run --race mutex.go\nHello Gophers!\nWelcome
  to Goworld!\n\n$ go run --race mutex.go\nWelcome to Goworld!\nHello Gophers!\n</pre></div>\n\n</pre>\n\n<p>The
  above example is a preventive measure for race conditions with a mutex Lock on the
  shared resource used in the go routines.</p>\n<p>Let's break down the code step
  by step. First, we initialize a few variables:</p>\n<ul>\n<li>\n<p><code>wg</code>
  as <code>sync.WaitGroup</code>:\nThe <code>wg</code> is a waitgroup that will be
  used for the synchronization of the go-routines by blocking the main function to
  exit</p>\n</li>\n<li>\n<p><code>mut</code> as <code>sync.Mutex</code>:\nThe <code>mut</code>
  is a structure that internally holds a few integer values indicating either a blocked
  or unblocked state. The <code>sync.Mutex</code> has two private fields wiz. <code>state</code>
  and <code>sema</code>, the state holds the mutex state either <code>0</code>(unlocked)
  or <code>1</code> as locked. The <code>sema</code> field is a <code>uint32</code>
  that is used for blocking and signaling, it acts as a semaphore to manage blocking
  and unblocking goroutines trying to acquire the mutex.\nThis will be used to acquire
  <code>Lock</code> and <code>Unlock</code> on the shared resource while writing the
  data to the file or appending the data to the resource.</p>\n</li>\n<li>\n<p><code>sharedBuffer</code>
  as <code>[]byte</code>:\nThe <code>sharedBuffer</code> is the actually shared resource
  that will be used to hold the strings for keeping track of the data written to the
  file. It will be the resource that would require to lock for mutating its value(appending
  to the slice) among the go routines.</p>\n</li>\n</ul>\n<p>We add <code>2</code>
  to the <code>wg</code> to indicate to wait for the two go-routines to complete,
  in the next line we call two go routines as the function <code>WriteToFile</code>.
  The <code>WriteToFile</code> is a function that takes in quite a few parameters
  namely, the filename, the content to write, the reference to the sharedBuffer, waitgroup,
  and the mutex.</p>\n<p>Inside the function <code>WriteToFile</code>:</p>\n<ul>\n<li>We
  first <code>defer</code> the <code>waitgroups</code> as <code>Done</code> i.e. to
  call the <code>wg.Done</code> method at the end of the function call.</li>\n<li>Typecast
  the <code>contents</code> from <code>string</code> as the <code>[]byte</code>.</li>\n<li>Acquire
  the <code>mutex.Lock()</code> i.e. to say &quot;The below operations should not
  be done concurrently, one at a time&quot;. We then append the <code>contents</code>
  to the <code>buffer</code> which is a pointer to the <code>sharedBuffer</code> in
  the main function, so essentially we are trying to mutate the <code>sharedBuffer</code>
  in this function.</li>\n<li>Open the file as <code>O_APPEND</code> and <code>O_WRONLY</code>
  to indicate that the file should be opened in append/writing mode. (We have observed
  this type of operation in the <a href=\"https://www.meetgor.com/golang-file-write/\">Golang:
  File Write</a> article)</li>\n<li>Use the <code>Write</code> method to write the
  slice of bytes(contents) into the file that we opened. We have a <code>defer</code>
  for the closing of the file.</li>\n</ul>\n<p>We obviously check for errors and print
  them, but it could be a <code>panic</code> or <code>log</code> depending on the
  type of operation the application is doing.\nSo that is the all operation we want
  to do, so we finally open the lock with <code>mutex.Unlock()</code> which will allow
  the other go routine if any to access the resource and proceed with the operations.</p>\n<h3>Read
  Write Mutual Exclusion Lock</h3>\n<p>The mutual exclusion lock is good if you have
  a write operation-heavy application. However, if we have read write operations in
  more or the less in same proportion(read heavy) we don't want to have readers getting
  blocked when other readers are accessing the resource since it is not a mutation
  process.</p>\n<p>We could allow many readers to read simultaneously. But for the
  writing operation, we want to block the readers/writers. The writer should be given
  preference first in case of a lock by other writers. This would prevent a writer
  from waiting for readers to complete. This is usually referred to as <strong>Read
  Write Mutual Exclusion</strong>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;time&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">reader</span><span class=\"p\">(</span><span class=\"nx\">id</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">mutex</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">Mutex</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"p\">!</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">TryLock</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Reader %d blocked!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Unlock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Reader %d: read count %d\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">id</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">writer</span><span
  class=\"p\">(</span><span class=\"nx\">id</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">increment</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">count</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"kt\">int</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">mutex</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">sync</span><span
  class=\"p\">.</span><span class=\"nx\">Mutex</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Done</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">mutex</span><span class=\"p\">.</span><span
  class=\"nx\">Unlock</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">mutex</span><span class=\"p\">.</span><span
  class=\"nx\">Lock</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"w\"> </span><span class=\"o\">+=</span><span class=\"w\"> </span><span class=\"nx\">increment</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Sleep</span><span class=\"p\">(</span><span
  class=\"mi\">5</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"w\">
  </span><span class=\"nx\">time</span><span class=\"p\">.</span><span class=\"nx\">Millisecond</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Writer %d: wrote count %d\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">1</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">mutex</span><span class=\"w\"> </span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">Mutex</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">readers</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">5</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">writers</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">3</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Add</span><span class=\"p\">(</span><span
  class=\"nx\">readers</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">i</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">;</span><span class=\"w\">
  </span><span class=\"nx\">i</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">readers</span><span class=\"p\">;</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"o\">&amp;</span><span class=\"nx\">count</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">wg</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">mutex</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Add</span><span class=\"p\">(</span><span class=\"nx\">writers</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">0</span><span
  class=\"p\">;</span><span class=\"w\"> </span><span class=\"nx\">i</span><span class=\"w\">
  </span><span class=\"p\">&lt;</span><span class=\"w\"> </span><span class=\"nx\">writers</span><span
  class=\"p\">;</span><span class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">writer</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">count</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">wg</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">mutex</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Wait</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run --race no-rwmutex.go \nReader <span class=\"m\">0</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">3</span> blocked!\nReader
  <span class=\"m\">2</span> blocked!\nReader <span class=\"m\">1</span> blocked!\nReader
  <span class=\"m\">4</span> blocked!\nWriter <span class=\"m\">0</span>: wrote count
  <span class=\"m\">2</span>\nWriter <span class=\"m\">1</span>: wrote count <span
  class=\"m\">3</span>\nWriter <span class=\"m\">2</span>: wrote count <span class=\"m\">4</span>\n\n$
  go run --race no-rwmutex.go \nReader <span class=\"m\">2</span> blocked!\nReader
  <span class=\"m\">0</span>: <span class=\"nb\">read</span> count <span class=\"m\">1</span>\nReader
  <span class=\"m\">1</span> blocked!\nReader <span class=\"m\">4</span> blocked!\nReader
  <span class=\"m\">3</span> blocked!\nWriter <span class=\"m\">2</span>: wrote count
  <span class=\"m\">2</span>\nWriter <span class=\"m\">0</span>: wrote count <span
  class=\"m\">3</span>\nWriter <span class=\"m\">1</span>: wrote count <span class=\"m\">4</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above example has a <code>reader</code> and a <code>writer</code> method, the <code>reader</code>
  method simply has to read the shared resource <code>count</code>. It acquires a
  <code>Mutex</code> lock before reading and unlocks it thereafter. Similarly, the
  <code>writer</code> function is used for incrementing the <code>count</code> shared
  resource.</p>\n<p>The <code>reader</code> method has a <a href=\"https://pkg.go.dev/sync#Mutex.TryLock\">TryLock</a>
  method that tries to acquire a mutex lock on the resource, if the resource is already
  locked, the function will return <code>false</code> and hence we can say that the
  reading is blocked(just for demonstration). And if the function <code>TryLock</code>
  returns <code>true</code>, it will acquire the <code>Lock</code>. We further <code>defer</code>
  the <code>Unlock</code> and access the <code>count</code> variable which is passed
  as a reference.</p>\n<p>The <code>writer</code> method is simply acquiring the <code>Lock</code>
  and incrementing the <code>counter</code> and thereafter <code>Unlock</code> is
  called with <code>defer</code>.</p>\n<p>In the above code:</p>\n<ul>\n<li>The <code>reader</code>
  and the <code>writer</code> both might be waiting for the lock to be released, however,
  for readers to wait for reading doesn't make sense.</li>\n<li>Because if you would
  want to just read a particular memory address, there shouldn't be any locks for
  one reader to wait for other readers to finish.</li>\n<li>However, for writing,
  there has to be a lock. The mutex lock will lock the resource irrespective of the
  <code>reader</code> or <code>writer</code>.</li>\n</ul>\n<p>This might not be as
  visible here, but it might be the reason, that all the <code>readers</code> are
  blocked from reading due to another reader's or writer's lock.</p>\n<p>The <code>sync</code>
  package has the <a href=\"https://pkg.go.dev/sync#RWMutex\">RWMutex</a> that does
  this exact same thing. It is almost similar to the <code>Mutex</code> however, it
  would allow concurrent reading operation and prefer writing operation before readers
  to prevent writer starvation.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;sync&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;time&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">reader</span><span class=\"p\">(</span><span class=\"nx\">id</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">mutex</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">RWMutex</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">RUnlock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"p\">!</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">TryRLock</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Reader %d blocked!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Reader %d: read count %d\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">writer</span><span class=\"p\">(</span><span class=\"nx\">id</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">increment</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"kt\">int</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">sync</span><span class=\"p\">.</span><span
  class=\"nx\">WaitGroup</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">mutex</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">RWMutex</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">wg</span><span class=\"p\">.</span><span class=\"nx\">Done</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Unlock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">mutex</span><span
  class=\"p\">.</span><span class=\"nx\">Lock</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"o\">*</span><span
  class=\"nx\">count</span><span class=\"w\"> </span><span class=\"o\">+=</span><span
  class=\"w\"> </span><span class=\"nx\">increment</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">time</span><span class=\"p\">.</span><span
  class=\"nx\">Sleep</span><span class=\"p\">(</span><span class=\"mi\">5</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"w\"> </span><span class=\"nx\">time</span><span
  class=\"p\">.</span><span class=\"nx\">Millisecond</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Writer %d: wrote count %d\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">id</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">count</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">count</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">1</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">wg</span><span class=\"w\"> </span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">WaitGroup</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">mutex</span><span class=\"w\"> </span><span
  class=\"nx\">sync</span><span class=\"p\">.</span><span class=\"nx\">RWMutex</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">readers</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">5</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">writers</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">3</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">wg</span><span
  class=\"p\">.</span><span class=\"nx\">Add</span><span class=\"p\">(</span><span
  class=\"nx\">readers</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">i</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">;</span><span class=\"w\">
  </span><span class=\"nx\">i</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">readers</span><span class=\"p\">;</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"o\">&amp;</span><span class=\"nx\">count</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">wg</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">mutex</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Add</span><span class=\"p\">(</span><span class=\"nx\">writers</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">0</span><span
  class=\"p\">;</span><span class=\"w\"> </span><span class=\"nx\">i</span><span class=\"w\">
  </span><span class=\"p\">&lt;</span><span class=\"w\"> </span><span class=\"nx\">writers</span><span
  class=\"p\">;</span><span class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">++</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">go</span><span class=\"w\"> </span><span class=\"nx\">writer</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">count</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">wg</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">mutex</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">wg</span><span class=\"p\">.</span><span
  class=\"nx\">Wait</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run --race rwmutex.go \nReader <span class=\"m\">0</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">3</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">1</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">2</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nReader <span class=\"m\">4</span>: <span class=\"nb\">read</span>
  count <span class=\"m\">1</span>\nWriter <span class=\"m\">1</span>: wrote count
  <span class=\"m\">2</span>\nWriter <span class=\"m\">0</span>: wrote count <span
  class=\"m\">3</span>\nWriter <span class=\"m\">2</span>: wrote count <span class=\"m\">4</span>\nWriter
  <span class=\"m\">3</span>: wrote count <span class=\"m\">5</span>\nWriter <span
  class=\"m\">4</span>: wrote count <span class=\"m\">6</span>\n\n$ go run --race
  rwmutex.go \nReader <span class=\"m\">1</span> blocked!\nReader <span class=\"m\">0</span>:
  <span class=\"nb\">read</span> count <span class=\"m\">1</span>\nReader <span class=\"m\">2</span>
  blocked!\nReader <span class=\"m\">3</span> blocked!\nReader <span class=\"m\">4</span>
  blocked!\nWriter <span class=\"m\">4</span>: wrote count <span class=\"m\">2</span>\nWriter
  <span class=\"m\">0</span>: wrote count <span class=\"m\">3</span>\nWriter <span
  class=\"m\">1</span>: wrote count <span class=\"m\">4</span>\nWriter <span class=\"m\">3</span>:
  wrote count <span class=\"m\">5</span>\nWriter <span class=\"m\">2</span>: wrote
  count <span class=\"m\">6</span>\n</pre></div>\n\n</pre>\n\n<p>In the modified example,
  all the logic remains the same, just the <code>sync.Mutex</code> is replaced with
  <code>sync.RWMutex</code>. Also for trying to acquire the lock in the <code>reader</code>
  method <a href=\"https://pkg.go.dev/sync#RWMutex.RLock\">TryRLock</a> is used instead
  of <a href=\"https://pkg.go.dev/sync#RWMutex.TryRLock\">TryLock</a> which will check
  if the existing lock acquired is of a reader or writer, if it is a reader, it will
  return <code>true</code>, else <code>false</code>. Also the <code>Unlock</code>
  is replaced with the <code>RUnlock</code> method for releasing the read lock. In
  the <code>writer</code> method, everything remains the same so the writer has to
  acquire the lock irrespective of whether the current lock is from the reader/writer,
  so it is a normal <code>Lock</code> and <code>Unlock</code>.</p>\n<p>In the above
  example, we can see all the read operations sometimes are executed instead of getting
  blocked. This is due to the <code>RWMutex</code> Lock and Unlock on the read operation/function.</p>\n<ul>\n<li>When
  one reader is reading, it can't block other readers.</li>\n<li>However with simple
  <code>Mutex</code>, the reader is even blocked when another reader is reading.</li>\n<li>For
  the write operation, it will be blocked as usual, so if a writer is performing a
  write operation and a reader/readers come in, they will be blocked, also if in the
  meantime, while the resource is locked, another writer comes in, the writer will
  be given preference instead of waiting for all readers to complete. This prevents
  writer starvation.</li>\n</ul>\n<p>You can see, that the readers are still blocked
  the majority of the time, however, they are blocked due to the writer locking the
  resource and not any other reader.</p>\n<p>There is a bit more suitable <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/go-routines/file-rw.go\">example</a>
  related to <code>file</code> reading and writing with <code>RWMutex</code>, make
  sure to check that out to get a more clear understanding of an actual case of using
  the <code>RWMutex</code>.</p>\n<h3>Channels</h3>\n<p>This is a big part and I would
  like to delve into this topic in a separate post. There are some patterns like <code>fan-in</code>,
  <code>fan-out</code>, <code>worker-pool</code>, <code>pub-sub</code>, etc. which
  are really common in web applications and backend systems. These patterns we shall
  explore in the next article.</p>\n<p>Channels are a way to provide a safe and idiomatic
  way for Goroutines to exchange data and coordinate their execution without resorting
  to low-level mechanisms like shared memory or explicit locking.</p>\n<p>That's it
  from the 30th part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/go-routines\">100
  days of Golang</a> repository.</p>\n<h3>References:</h3>\n<ul>\n<li><a href=\"https://blog.nindalf.com/posts/how-goroutines-work/\">How
  Goroutines Work</a></li>\n<li><a href=\"https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/\">Concurrency
  patterns in Golang: WaitGroups and Goroutines</a></li>\n<li><a href=\"https://riteeksrivastava.medium.com/a-complete-journey-with-goroutines-8472630c7f5c\">A
  complete journey with Goroutines</a></li>\n</ul>\n<h2>Conclusion</h2>\n<p>From this
  part of the series, the fundamentals of golang's concurrency model were understood
  specifically spawning go-routines, synchronously executing go-routines with the
  help of a wait group, mutex locks, and how to secure concurrent access to a shared
  resource. In the next part of the series, these concepts will be used in asynchronous
  communication using channels.</p>\n<p>Hopefully, you have got the basics of concurrency
  in golang cleared from this post. If you have any queries, suggestions, or feedback,
  please feel free to comment below or contact me on the social handles. Thank you
  for reading. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/golang-030-goroutines.png
long_description: One of the key features that set Go apart from many other languages
  is its native support for Goroutines - lightweight concurrent functions that can
  run concurrently and efficiently manage concurrency tasks. Goroutines are an essential
  aspect of Go I
now: 2025-05-01 18:11:33.313866
path: blog/posts/2023-07-29-GO-30-GoRoutines.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-go-routines
status: published
tags:
- go
templateKey: blog-post
title: 'Golang: Go Routines and WaitGroups'
today: 2025-05-01
---

## Introduction

One of the key features that set Go apart from many other languages is its native support for Goroutines - lightweight concurrent functions that can run concurrently and efficiently manage concurrency tasks. Goroutines are an essential aspect of Go's concurrency model, enabling developers to build highly concurrent and performant applications effortlessly.

In this blog post, we will dive into the world of Goroutines and explore how they work, why they are essential for concurrent programming in Go, and how they can greatly improve the responsiveness and efficiency of your applications. This post will cover go routines and a primer on wait groups, in the next article we will be looking deeply into channels where all these three things can be demonstrated and understood in a better and more useful way.

## Concurrency and Parallelism

This two concepts are quite crucial before diving into the fundamentals of concurreny.

**Concurrency** is about dealing with multiple things at once. Concurrent programs can have several processes or threads running simultaneously on a single CPU core by rapidly switching between them (context switching). The threads are interleaved, not necessarily executing at literally the same time. The CPU can switch between these tasks to give the appearance of simultaneous progress.

**Parallelism** is about doing multiple things at literally the same time. Parallel programs can execute multiple computations simultaneously on separate CPU cores. The threads actually execute in parallel.

## What is a Go Routine

A go routine is a simple lightweight thread managed by the Go runtime. In the simplest of terms, a go routine can be defined as:

> Go routine is a way to perform multiple tasks within a program, allowing different parts of the program to work simultaneously and make the most out of the resources.

Also it can be stated as:

A goroutine in Golang is a lightweight, independently executing function that runs concurrently with other goroutines within the same address space. In other words, it is a concurrent unit of execution.

Focussing on the word **same address space** that will be really critical in the later sections of this article.

## Features of Go Routines

Go routines form a key part of Go's concurrency model. Here are some of the key features of go routines:

- **Lightweight Thread**:
    A Go routine is often referred to as a lightweight thread.

- **Independent Execution**:
    Go routines run independently of each other, enabling concurrent execution.

- **Managed by Go**:
    Go routines are managed by the Go runtime, making them easy to use.

- **Low Overhead**:
    Go routines have low memory overhead, allowing us to create thousands of them efficiently.

- **Communication**:
    Go routines can communicate and synchronize data through channels.

- **Asynchronous**:
    Go routines can execute asynchronously, allowing other parts of the program to continue running.

- **Scalability**:
    Go routines are the foundation of scalable concurrent programming in Go.

Unlike threads in other languages, goroutines are cheap and you can easily create thousands or even millions of them in a program.


## Example of Go Routines

Creating a go routine is not hard, just add the keyword `go` before the function call and the go runtime will create a new go routine inside the main function or wherever is the context from. Remember, the main function is also a go routine.

```go
package main

import (
	"fmt"
	"time"
)

func process() {
	fmt.Println("Hello World!")
	time.Sleep(time.Second)
}

func main() {
	start := time.Now()
	go process()
	go process()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println(duration)
}
```

```bash
scripts/go-routines on  main via 🐹 v1.20 
$ go run main.go 

Hello World!
15.607µs


scripts/go-routines on  main via 🐹 v1.20 
$ go run main.go 

9.889µs


scripts/go-routines on  main via 🐹 v1.20 
$ go run main.go 

8.834µs


scripts/go-routines on  main via 🐹 v1.20 
$ go run main.go 

9.158µs

scripts/go-routines on  main via 🐹 v1.20 
$ go run main.go 

12.54µs


scripts/go-routines on  main via 🐹 v1.20 
$ go run main.go 

Hello World!
Hello World!
10.19µs


scripts/go-routines on  main via 🐹 v1.20 
$ go run main.go 

14.1µs


scripts/go-routines on  main via 🐹 v1.20 
$ 
```

Quite unpredictable output right? This is a power of go-routines. It is **asynchronous** so it will not block the main function. The two function calls to `process`, are executed independently of the main function scope. The program just calls the process function and captures the output, and sequentially reaches the end of the program(main function), At this point the go routines(threads) inside the main function are stopped abruptly.

Let's break it down if it is not clear yet.

- The main function starts.
- `go process()` creates a Go routine and starts its execution.
- Meanwhile, another call to `go process()` creates a separate Go routine and starts its execution as well.
- Meanwhile, calculate the time difference between the main function start and end.
- Meanwhile main function ends.

So, in summary, the main function is only able to capture the output of the `duration` as it is synchronous and if the process has been executed, it prints the `Hello World!` message. Hence we have different outputs because the uncontrolled concurrency, lack of coordination, and OS scheduling across the program runs are different. 



## Wait Groups

In simple terms, WaitGroup is used to synchronize multiple goroutines and to wait for them to finish executing. This allows the go routines to be completed before the completion of the main function, hence it blocks the main function from leaving/exiting the scope.

> A WaitGroup is a synchronization primitive that allows a goroutine to wait for a collection of other goroutines to finish executing. 

- A WaitGroup is initialized with a counter representing the number of goroutines to wait for.
- The Add() method increments the counter by the given value. This is called by each goroutine to indicate it is running.
- The main goroutine calls Add() to set the initial count, then launches worker goroutines.
- A WaitGroup is typically passed by a pointer to goroutines that need to be waited on.
- The Done() method decrements the counter by 1. Goroutines call this when finished.
- Each worker calls Done() when finished, decrementing the counter.
- The Wait() method blocks until the counter reaches 0, indicating all goroutines have finished.
- Main calls Wait() to block until Done() brings counter to 0.

This provides a simple way to synchronize multiple goroutines finishing their work with a main thread that needs to wait for them to complete. The counter ensures the main thread knows how many goroutines it is waiting for. Interacting and working with go routines with synchronization using a wait group is quite intuitive and simple to follow. Let's look at a simple example below:


```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func process(pid int, wg *sync.WaitGroup) {
	fmt.Printf("Started process %d\n", pid)
	time.Sleep(1 * time.Second)
	fmt.Printf("Completed process %d\n", pid)
	defer wg.Done()
}

func main() {

	now := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All processes completed")
	end := time.Now()
	fmt.Println(end.Sub(now))
}
```

```bash
scripts/go-routines on  main [?] via 🐹 v1.20 
$ go run wg.go 
Started process 9
Started process 0
Started process 1
Started process 2
Started process 3
Started process 4
Started process 5
Started process 6
Started process 7
Started process 8
Completed process 8
Completed process 3
Completed process 9
Completed process 0
Completed process 1
Completed process 2
Completed process 5
Completed process 4
Completed process 6
Completed process 7
1.000563257s
All processes completed

scripts/go-routines on  main [?] via 🐹 v1.20 
$ 

```

In the above example, we have used the same function `process` but with a slight twist, we have added a process id, just an integer to represent the go routine. We print the start and completion of the function in between the sleep for 1 second. We also have a wait group. 
- A [WaitGroup](https://pkg.go.dev/sync#WaitGroup) is basically a struct type defined in the [sync](https://pkg.go.dev/sync) package. 
- The variable `wg` is a new wait group instance that will be used to synchronize and wait for the completion of groups of go routines. 

We create a for loop with `10` iterations, so as to spawn 10 `process` function calls. We first use the `wg.Add(1)` which says to the wait group to wait for 1 go routine. The immediate next line is a go routine `go process()` which takes in a `pid` just to track which go routine is being executed in the loop. 

Inside the `process` function, we simply say the process with the given `pid` has started as a print statement, sleeps for a second, and then print the end of the process with `pid`. The end of the function is marked with `defer wg.Done()`, this is to indicate that the go routine is completed. 

The wg(wait group) has a counter that keeps track of the number of go routines it has to synchronize or wait for them to complete. In the [Add](https://pkg.go.dev/sync#WaitGroup.Add) function, the internal counter in the `WaitGroup` is incremented by the `delta` the integer parsed as a parameter. And the [Done](https://pkg.go.dev/sync#WaitGroup.Done) function, decrements the internal counter in the `WaitGroup` which indicates that the go routine is completed.

In the main function, we call the [wg.Wait](https://pkg.go.dev/sync#WaitGroup.Wait) which will block until the counter for the `WaitGroup` is 0 i.e. all go routines have completed their execution. So, we created 10 go routines which ran concurrently but were synchronized with the help of WaitGroups. The WaitGroup is allowed to block the main function till all the go routines are done executing.


## Go Routine with WaitGroup Use Cases

Go routines can be used for creating asynchronous tasks, and also for creating concurrent tasks. By using wait groups, we can create a way to wait for multiple goroutines to complete. By using go routines and wait groups, we can complete n number of tasks in a time frame of 1 task's completion. However, to create concurrent communication between the other processes, we need `channels` (which we will explore in the next article). 

Here is a simple breakdown of what asynchronous and concurrent tasks might refer to:

**Asynchronous** tasks run independently of the main program flow, allowing the main program to continue executing without waiting for the task's completion. For example, not blocking other tasks in the sequential flow of the main function.

**Concurrent** tasks run simultaneously and can execute at the same time as other tasks. They make use of multiple threads (goroutines in Go's case) to achieve parallel execution. For example, running multiple tasks parallel will cut the time down for spinning each task after the completion of another task.

Some of the asynchronous tasks that can be done with go routines might include the following:

- Sending a mail while saving the user to the database.
- Fetch and process data from multiple websites(web scraping/crawling).
- High-performance message brokers and queuing systems for inter-process communication.


One practical way of using go routines with wait groups would be to send mail, we are not going to see the actual implementation of the mail-sending stuff. However, we can try to mimic how the setup might be for sending bulk mail. By creating a wait group and a list of mail addresses to send, a go routine can be created with a function that sends those emails.

```go
package main

import (
	"fmt"
	"sync"
)

func sendMail(address string, wg *sync.WaitGroup) {
	fmt.Println("Sending mail to", address)
	defer wg.Done()
    // Actual mail sending, smtp stuff
    // handle errors

    // client, err := smtp.Dial("smtp.example.com:587")
    // errr = client.Mail("sender@example.com")
    // err = client.Rcpt(address)
    // wc, err := client.Data()
    //_, err = wc.Write([]byte("This is the email body."))
    // err = wc.Close()
    // client.Quit()
}

func main() {
	emails := []string{
		"recipient1@example.com",
		"recipient2@example.com",
		"xyz@example.com",
	}
	wg := sync.WaitGroup{}
    wg.Add(len(emails))

	for _, email := range emails {
		go sendMail(email, &wg)
	}
	wg.Wait()
	fmt.Println("All emails queued for sending")
	// Do other stuff
}
```
In the above example, the `emails` is a list of email ids to send the mail. We have created a wait group that we can initialize with the total number of go routines probably to be executed. The `wg.Add` method is parsed with the number of emails to be sent hence that equals the go routines to spawn.

So, we can then in the for loop, iterate over each mail and send the emails with the `sendMail` function as go routine. The `wg.Wait` function outside the loop, will make sure the main function is halted before all the go routines complete their execution.

There is one more way to call a function as a go routine without changing its signature since we had to pass `wg` pointer as the WaitGroup reference to it to acknowledge the completion of the go routine. We could wrap these two operations wiz. calling the function and calling the `wg.Done` method with an anonymous function.

```go
package main

import (
	"fmt"
	"sync"
)

func sendMail(address string) {
	fmt.Println("Sending mail to", address)
}

func main() {
	emails := []string{
		"recipient1@example.com",
		"recipient2@example.com",
		"xyz@example.com",
	}
	wg := sync.WaitGroup{}
	wg.Add(len(emails))

	for _, email := range emails {
		mail := email
		go func(m string) {
			sendMail(m)
			wg.Done()
		}(mail)
	}
	wg.Wait()
	fmt.Println("All emails queued for sending")
	// Do other stuff
}
```

This does the exact same thing, but we don't have to change the signature of the function, This keeps the functional logic of the function in its place and handles the concurrency on the go.

```go

for _, email := range emails {
    mail := email
    go func(m string) {
        sendMail(m)
        wg.Done()
    }(mail)
}
```

If the above bit scares you, don't worry it's too simple. 

- We are iterating over the email slice using a for loop and creating a Go routine for each email address. The loop variable email represents the email address at the current iteration.
- However, to avoid the loop variable capture issue (where all Go routines would share the same email variable), we create a new variable mail and assign the value of email to it. This step ensures that each Go routine captures its own copy of the email address.
- We immediately create an anonymous function (a closure) using the go keyword. This anonymous function takes the `mail` variable as a parameter m and is executed concurrently as a Go routine. Inside the Go routine, we call the `sendMail` function with the email address `m`.
- After the `sendMail` call has been executed i.e. email has been sent, we call wg.Done() to notify the WaitGroup that the Go routine has completed its work. This allows the WaitGroup to properly synchronize and wait for all Go routines to finish before the program proceeds beyond wg.Wait() in the main function.

This is one way to do it if you want to separate the mail-sending logic from the goroutines/concurrency task. However, this should be handled with care as the variables inside the closure might be shared among all the goroutines instead of having individual variable literals. 

To ensure that each goroutine operates on its own copy of the email address, we use the approach of creating a new variable mail and passing it as a parameter to the anonymous function. This way, each goroutine captures its unique email address, avoiding any interference or unintended sharing of data between goroutines.

## Mutual Exclusion Locks

In the previous examples, we saw how goroutines and wait groups allow us to run multiple tasks concurrently in Go. However, sometimes these concurrent goroutines need to access shared resources like memory, files, network sockets, etc.

When more than one goroutine tries to access a resource at the same time, it can lead to **race conditions** and unpredictable behavior. To handle this, we need a way to ensure only one goroutine can access the resource at a time.

This is where mutual exclusion locks come in. A **mutual exclusion lock**, or **mutex**, provides a mechanism to lock access to a shared resource. It ensures only one goroutine at a time can acquire the lock, blocking other goroutines until the lock is released.

For example, say we have multiple goroutines trying to append data to the same memory buffer(could be file/database/etc.) concurrently:

```go
package main

import (
	"fmt"
	"os"
	"sync"
)

func WriteToFile(filename string, contents string, buffer *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	*buffer = append(*buffer, []byte(contents)...)
	err := os.WriteFile(filename, *buffer, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var wg sync.WaitGroup
	var sharedBuffer []byte

	wg.Add(2)
	go WriteToFile("data/f1.txt", "Hello ", &sharedBuffer, &wg)
	go WriteToFile("data/f1.txt", "World! ", &sharedBuffer, &wg)
	wg.Wait()

	fmt.Println(string(sharedBuffer))
}
```

```bash
$ go run --race no-mutex.go
==================
WARNING: DATA RACE
Read at 0x00c000012030 by goroutine 8:
  main.WriteToFile()
      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:11 +0xe7
  main.main.func2()
      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:24 +0x64

Previous write at 0x00c000012030 by goroutine 7:
  main.WriteToFile()
      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:11 +0x16a
  main.main.func1()
      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:23 +0x64

Goroutine 8 (running) created at:
  main.main()
      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:24 +0x1f6

Goroutine 7 (finished) created at:
  main.main()
      /home/meet/code/100-days-of-golang/scripts/go-routines/no-mutex.go:23 +0x14e
==================
```

This is because we are trying to access the same memory address at the same time. This is a race condition and can lead to unpredictable behavior. Try removing the `--race` flag while running, in this little stupid example, it might not be obvious, but in complex and constrained environments, this can get the application in serious trouble.

**NOTE**: We are using `go run --race no-mutex.go` to check if there are any race conditions in the program. This is the [race detector](https://go.dev/doc/articles/race_detector) flag in the run command.

To avoid this race condition, we need to add the mutex locks provided in the [sync.Mutex](https://pkg.go.dev/sync#Mutex) type. There are methods like [Lock](), [Unlock](), and [TryLock]() which help in locking access of the resource to a single entity at a given time.

When a goroutine calls `Lock()` on a mutex, it acquires the lock. If the mutex is already locked by another goroutine, the calling goroutine will be blocked (put to sleep) until the lock becomes available. Once the lock is acquired successfully, the goroutine can proceed with its critical section, which is the part of the code that should not be executed concurrently by multiple goroutines.

When a goroutine calls `Unlock()` on a mutex, it releases the lock. This allows other waiting goroutines to acquire the lock and execute their critical sections. It's essential to ensure that `Unlock()` is called after the critical section to release the mutex and avoid deadlocks. The critical section/shared resource should not be accessed after the release of this lock.

```go
package main

import (
	"fmt"
	"os"
	"sync"
)

func WriteToFile(filename string, contents string, buffer *[]byte, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	contentBytes := []byte(contents)

	mutex.Lock()
	*buffer = append(*buffer, contentBytes...)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.Write(contentBytes)
	if err != nil {
		fmt.Println(err)
	}
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	var sharedBuffer []byte

	wg.Add(2)
	go WriteToFile("data/f1.txt", "Hello Gophers!\n", &sharedBuffer, &wg, &mut)
	go WriteToFile("data/f1.txt", "Welcome to Goworld!\n", &sharedBuffer, &wg, &mut)
	wg.Wait()

	fmt.Println(string(sharedBuffer))
}
```

```bash
$ go run --race mutex.go
Welcome to Goworld!
Hello Gophers!

$ go run --race mutex.go
Hello Gophers!
Welcome to Goworld!

$ go run --race mutex.go
Hello Gophers!
Welcome to Goworld!

$ go run --race mutex.go
Welcome to Goworld!
Hello Gophers!
```

The above example is a preventive measure for race conditions with a mutex Lock on the shared resource used in the go routines. 

Let's break down the code step by step. First, we initialize a few variables:
- `wg` as `sync.WaitGroup`:
    The `wg` is a waitgroup that will be used for the synchronization of the go-routines by blocking the main function to exit

- `mut` as `sync.Mutex`:
    The `mut` is a structure that internally holds a few integer values indicating either a blocked or unblocked state. The `sync.Mutex` has two private fields wiz. `state` and `sema`, the state holds the mutex state either `0`(unlocked) or `1` as locked. The `sema` field is a `uint32` that is used for blocking and signaling, it acts as a semaphore to manage blocking and unblocking goroutines trying to acquire the mutex.
    This will be used to acquire `Lock` and `Unlock` on the shared resource while writing the data to the file or appending the data to the resource.

- `sharedBuffer` as `[]byte`: 
  The `sharedBuffer` is the actually shared resource that will be used to hold the strings for keeping track of the data written to the file. It will be the resource that would require to lock for mutating its value(appending to the slice) among the go routines.

We add `2` to the `wg` to indicate to wait for the two go-routines to complete, in the next line we call two go routines as the function `WriteToFile`. The `WriteToFile` is a function that takes in quite a few parameters namely, the filename, the content to write, the reference to the sharedBuffer, waitgroup, and the mutex.

Inside the function `WriteToFile`:
-  We first `defer` the `waitgroups` as `Done` i.e. to call the `wg.Done` method at the end of the function call. 
- Typecast the `contents` from `string` as the `[]byte`. 
- Acquire the `mutex.Lock()` i.e. to say "The below operations should not be done concurrently, one at a time". We then append the `contents` to the `buffer` which is a pointer to the `sharedBuffer` in the main function, so essentially we are trying to mutate the `sharedBuffer` in this function.
- Open the file as `O_APPEND` and `O_WRONLY` to indicate that the file should be opened in append/writing mode. (We have observed this type of operation in the [Golang: File Write](https://www.meetgor.com/golang-file-write/) article)
- Use the `Write` method to write the slice of bytes(contents) into the file that we opened. We have a `defer` for the closing of the file. 

We obviously check for errors and print them, but it could be a `panic` or `log` depending on the type of operation the application is doing.
So that is the all operation we want to do, so we finally open the lock with `mutex.Unlock()` which will allow the other go routine if any to access the resource and proceed with the operations.

### Read Write Mutual Exclusion Lock

The mutual exclusion lock is good if you have a write operation-heavy application. However, if we have read write operations in more or the less in same proportion(read heavy) we don't want to have readers getting blocked when other readers are accessing the resource since it is not a mutation process. 

We could allow many readers to read simultaneously. But for the writing operation, we want to block the readers/writers. The writer should be given preference first in case of a lock by other writers. This would prevent a writer from waiting for readers to complete. This is usually referred to as **Read Write Mutual Exclusion**.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func reader(id int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	if !mutex.TryLock() {
		fmt.Printf("Reader %d blocked!\n", id)
		return
	}
	defer mutex.Unlock()
	fmt.Printf("Reader %d: read count %d\n", id, *count)
}

func writer(id, increment int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	*count += increment
	time.Sleep(5 * time.Millisecond)
	fmt.Printf("Writer %d: wrote count %d\n", id, *count)
}

func main() {

	count := 1
	var wg sync.WaitGroup
	var mutex sync.Mutex
    readers := 5
    writers := 3
	wg.Add(readers)
	for i := 0; i < readers; i++ {
		go reader(i, &count, &wg, &mutex)
	}

	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go writer(i, 1, &count, &wg, &mutex)
	}
	wg.Wait()

}
```

```bash
$ go run --race no-rwmutex.go 
Reader 0: read count 1
Reader 3 blocked!
Reader 2 blocked!
Reader 1 blocked!
Reader 4 blocked!
Writer 0: wrote count 2
Writer 1: wrote count 3
Writer 2: wrote count 4

$ go run --race no-rwmutex.go 
Reader 2 blocked!
Reader 0: read count 1
Reader 1 blocked!
Reader 4 blocked!
Reader 3 blocked!
Writer 2: wrote count 2
Writer 0: wrote count 3
Writer 1: wrote count 4
```

The above example has a `reader` and a `writer` method, the `reader` method simply has to read the shared resource `count`. It acquires a `Mutex` lock before reading and unlocks it thereafter. Similarly, the `writer` function is used for incrementing the `count` shared resource. 

The `reader` method has a [TryLock](https://pkg.go.dev/sync#Mutex.TryLock) method that tries to acquire a mutex lock on the resource, if the resource is already locked, the function will return `false` and hence we can say that the reading is blocked(just for demonstration). And if the function `TryLock` returns `true`, it will acquire the `Lock`. We further `defer` the `Unlock` and access the `count` variable which is passed as a reference. 

The `writer` method is simply acquiring the `Lock` and incrementing the `counter` and thereafter `Unlock` is called with `defer`.

In the above code:
- The `reader` and the `writer` both might be waiting for the lock to be released, however, for readers to wait for reading doesn't make sense. 
- Because if you would want to just read a particular memory address, there shouldn't be any locks for one reader to wait for other readers to finish.
- However, for writing, there has to be a lock. The mutex lock will lock the resource irrespective of the `reader` or `writer`. 

This might not be as visible here, but it might be the reason, that all the `readers` are blocked from reading due to another reader's or writer's lock.

The `sync` package has the [RWMutex](https://pkg.go.dev/sync#RWMutex) that does this exact same thing. It is almost similar to the `Mutex` however, it would allow concurrent reading operation and prefer writing operation before readers to prevent writer starvation.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func reader(id int, count *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	defer mutex.RUnlock()

	if !mutex.TryRLock() {
		fmt.Printf("Reader %d blocked!\n", id)
		return
	}
	fmt.Printf("Reader %d: read count %d\n", id, *count)

}

func writer(id, increment int, count *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()

	*count += increment
	time.Sleep(5 * time.Millisecond)
	fmt.Printf("Writer %d: wrote count %d\n", id, *count)
}

func main() {

	count := 1
	var wg sync.WaitGroup
	var mutex sync.RWMutex
    readers := 5
    writers := 3
	wg.Add(readers)
	for i := 0; i < readers; i++ {
		go reader(i, &count, &wg, &mutex)
	}

	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go writer(i, 1, &count, &wg, &mutex)
	}
	wg.Wait()

}
```

```bash
$ go run --race rwmutex.go 
Reader 0: read count 1
Reader 3: read count 1
Reader 1: read count 1
Reader 2: read count 1
Reader 4: read count 1
Writer 1: wrote count 2
Writer 0: wrote count 3
Writer 2: wrote count 4
Writer 3: wrote count 5
Writer 4: wrote count 6

$ go run --race rwmutex.go 
Reader 1 blocked!
Reader 0: read count 1
Reader 2 blocked!
Reader 3 blocked!
Reader 4 blocked!
Writer 4: wrote count 2
Writer 0: wrote count 3
Writer 1: wrote count 4
Writer 3: wrote count 5
Writer 2: wrote count 6

```

In the modified example, all the logic remains the same, just the `sync.Mutex` is replaced with `sync.RWMutex`. Also for trying to acquire the lock in the `reader` method [TryRLock](https://pkg.go.dev/sync#RWMutex.RLock) is used instead of [TryLock](https://pkg.go.dev/sync#RWMutex.TryRLock) which will check if the existing lock acquired is of a reader or writer, if it is a reader, it will return `true`, else `false`. Also the `Unlock` is replaced with the `RUnlock` method for releasing the read lock. In the `writer` method, everything remains the same so the writer has to acquire the lock irrespective of whether the current lock is from the reader/writer, so it is a normal `Lock` and `Unlock`.

In the above example, we can see all the read operations sometimes are executed instead of getting blocked. This is due to the `RWMutex` Lock and Unlock on the read operation/function. 
- When one reader is reading, it can't block other readers. 
- However with simple `Mutex`, the reader is even blocked when another reader is reading.
- For the write operation, it will be blocked as usual, so if a writer is performing a write operation and a reader/readers come in, they will be blocked, also if in the meantime, while the resource is locked, another writer comes in, the writer will be given preference instead of waiting for all readers to complete. This prevents writer starvation.

You can see, that the readers are still blocked the majority of the time, however, they are blocked due to the writer locking the resource and not any other reader.

There is a bit more suitable [example](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/go-routines/file-rw.go) related to `file` reading and writing with `RWMutex`, make sure to check that out to get a more clear understanding of an actual case of using the `RWMutex`.

### Channels

This is a big part and I would like to delve into this topic in a separate post. There are some patterns like `fan-in`, `fan-out`, `worker-pool`, `pub-sub`, etc. which are really common in web applications and backend systems. These patterns we shall explore in the next article.

Channels are a way to provide a safe and idiomatic way for Goroutines to exchange data and coordinate their execution without resorting to low-level mechanisms like shared memory or explicit locking.

That's it from the 30th part of the series, all the source code for the examples are linked in the GitHub on the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/go-routines) repository.

### References:

- [How Goroutines Work](https://blog.nindalf.com/posts/how-goroutines-work/)
- [Concurrency patterns in Golang: WaitGroups and Goroutines](https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/)
- [A complete journey with Goroutines](https://riteeksrivastava.medium.com/a-complete-journey-with-goroutines-8472630c7f5c)

## Conclusion

From this part of the series, the fundamentals of golang's concurrency model were understood specifically spawning go-routines, synchronously executing go-routines with the help of a wait group, mutex locks, and how to secure concurrent access to a shared resource. In the next part of the series, these concepts will be used in asynchronous communication using channels. 

Hopefully, you have got the basics of concurrency in golang cleared from this post. If you have any queries, suggestions, or feedback, please feel free to comment below or contact me on the social handles. Thank you for reading. Happy Coding :)