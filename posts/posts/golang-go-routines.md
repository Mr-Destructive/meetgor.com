{
  "title": "Golang: Go Routines and WaitGroups",
  "status": "published",
  "slug": "golang-go-routines",
  "date": "2025-04-05T12:33:30Z"
}

<h2>Introduction</h2>
<p>One of the key features that set Go apart from many other languages is its native support for Goroutines - lightweight concurrent functions that can run concurrently and efficiently manage concurrency tasks. Goroutines are an essential aspect of Go's concurrency model, enabling developers to build highly concurrent and performant applications effortlessly.</p>
<p>In this blog post, we will dive into the world of Goroutines and explore how they work, why they are essential for concurrent programming in Go, and how they can greatly improve the responsiveness and efficiency of your applications. This post will cover go routines and a primer on wait groups, in the next article we will be looking deeply into channels where all these three things can be demonstrated and understood in a better and more useful way.</p>
<h2>Concurrency and Parallelism</h2>
<p>This two concepts are quite crucial before diving into the fundamentals of concurreny.</p>
<p><strong>Concurrency</strong> is about dealing with multiple things at once. Concurrent programs can have several processes or threads running simultaneously on a single CPU core by rapidly switching between them (context switching). The threads are interleaved, not necessarily executing at literally the same time. The CPU can switch between these tasks to give the appearance of simultaneous progress.</p>
<p><strong>Parallelism</strong> is about doing multiple things at literally the same time. Parallel programs can execute multiple computations simultaneously on separate CPU cores. The threads actually execute in parallel.</p>
<h2>What is a Go Routine</h2>
<p>A go routine is a simple lightweight thread managed by the Go runtime. In the simplest of terms, a go routine can be defined as:</p>
<blockquote>
<p>Go routine is a way to perform multiple tasks within a program, allowing different parts of the program to work simultaneously and make the most out of the resources.</p>
</blockquote>
<p>Also it can be stated as:</p>
<p>A goroutine in Golang is a lightweight, independently executing function that runs concurrently with other goroutines within the same address space. In other words, it is a concurrent unit of execution.</p>
<p>Focussing on the word <strong>same address space</strong> that will be really critical in the later sections of this article.</p>
<h2>Features of Go Routines</h2>
<p>Go routines form a key part of Go's concurrency model. Here are some of the key features of go routines:</p>
<ul>
<li>
<p><strong>Lightweight Thread</strong>:
A Go routine is often referred to as a lightweight thread.</p>
</li>
<li>
<p><strong>Independent Execution</strong>:
Go routines run independently of each other, enabling concurrent execution.</p>
</li>
<li>
<p><strong>Managed by Go</strong>:
Go routines are managed by the Go runtime, making them easy to use.</p>
</li>
<li>
<p><strong>Low Overhead</strong>:
Go routines have low memory overhead, allowing us to create thousands of them efficiently.</p>
</li>
<li>
<p><strong>Communication</strong>:
Go routines can communicate and synchronize data through channels.</p>
</li>
<li>
<p><strong>Asynchronous</strong>:
Go routines can execute asynchronously, allowing other parts of the program to continue running.</p>
</li>
<li>
<p><strong>Scalability</strong>:
Go routines are the foundation of scalable concurrent programming in Go.</p>
</li>
</ul>
<p>Unlike threads in other languages, goroutines are cheap and you can easily create thousands or even millions of them in a program.</p>
<h2>Example of Go Routines</h2>
<p>Creating a go routine is not hard, just add the keyword <code>go</code> before the function call and the go runtime will create a new go routine inside the main function or wherever is the context from. Remember, the main function is also a go routine.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;time&quot;
)

func process() {
	fmt.Println(&quot;Hello World!&quot;)
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
</code></pre>
<pre><code class="language-bash">scripts/go-routines on  main via 🐹 v1.20 
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
</code></pre>
<p>Quite unpredictable output right? This is a power of go-routines. It is <strong>asynchronous</strong> so it will not block the main function. The two function calls to <code>process</code>, are executed independently of the main function scope. The program just calls the process function and captures the output, and sequentially reaches the end of the program(main function), At this point the go routines(threads) inside the main function are stopped abruptly.</p>
<p>Let's break it down if it is not clear yet.</p>
<ul>
<li>The main function starts.</li>
<li><code>go process()</code> creates a Go routine and starts its execution.</li>
<li>Meanwhile, another call to <code>go process()</code> creates a separate Go routine and starts its execution as well.</li>
<li>Meanwhile, calculate the time difference between the main function start and end.</li>
<li>Meanwhile main function ends.</li>
</ul>
<p>So, in summary, the main function is only able to capture the output of the <code>duration</code> as it is synchronous and if the process has been executed, it prints the <code>Hello World!</code> message. Hence we have different outputs because the uncontrolled concurrency, lack of coordination, and OS scheduling across the program runs are different.</p>
<h2>Wait Groups</h2>
<p>In simple terms, WaitGroup is used to synchronize multiple goroutines and to wait for them to finish executing. This allows the go routines to be completed before the completion of the main function, hence it blocks the main function from leaving/exiting the scope.</p>
<blockquote>
<p>A WaitGroup is a synchronization primitive that allows a goroutine to wait for a collection of other goroutines to finish executing.</p>
</blockquote>
<ul>
<li>A WaitGroup is initialized with a counter representing the number of goroutines to wait for.</li>
<li>The Add() method increments the counter by the given value. This is called by each goroutine to indicate it is running.</li>
<li>The main goroutine calls Add() to set the initial count, then launches worker goroutines.</li>
<li>A WaitGroup is typically passed by a pointer to goroutines that need to be waited on.</li>
<li>The Done() method decrements the counter by 1. Goroutines call this when finished.</li>
<li>Each worker calls Done() when finished, decrementing the counter.</li>
<li>The Wait() method blocks until the counter reaches 0, indicating all goroutines have finished.</li>
<li>Main calls Wait() to block until Done() brings counter to 0.</li>
</ul>
<p>This provides a simple way to synchronize multiple goroutines finishing their work with a main thread that needs to wait for them to complete. The counter ensures the main thread knows how many goroutines it is waiting for. Interacting and working with go routines with synchronization using a wait group is quite intuitive and simple to follow. Let's look at a simple example below:</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
	&quot;time&quot;
)

func process(pid int, wg *sync.WaitGroup) {
	fmt.Printf(&quot;Started process %d
&quot;, pid)
	time.Sleep(1 * time.Second)
	fmt.Printf(&quot;Completed process %d
&quot;, pid)
	defer wg.Done()
}

func main() {

	now := time.Now()
	var wg sync.WaitGroup

	for i := 0; i &lt; 10; i++ {
		wg.Add(1)
		go process(i, &amp;wg)
	}
	wg.Wait()
	fmt.Println(&quot;All processes completed&quot;)
	end := time.Now()
	fmt.Println(end.Sub(now))
}
</code></pre>
<pre><code class="language-bash">scripts/go-routines on  main [?] via 🐹 v1.20 
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

</code></pre>
<p>In the above example, we have used the same function <code>process</code> but with a slight twist, we have added a process id, just an integer to represent the go routine. We print the start and completion of the function in between the sleep for 1 second. We also have a wait group.</p>
<ul>
<li>A <a href="https://pkg.go.dev/sync#WaitGroup">WaitGroup</a> is basically a struct type defined in the <a href="https://pkg.go.dev/sync">sync</a> package.</li>
<li>The variable <code>wg</code> is a new wait group instance that will be used to synchronize and wait for the completion of groups of go routines.</li>
</ul>
<p>We create a for loop with <code>10</code> iterations, so as to spawn 10 <code>process</code> function calls. We first use the <code>wg.Add(1)</code> which says to the wait group to wait for 1 go routine. The immediate next line is a go routine <code>go process()</code> which takes in a <code>pid</code> just to track which go routine is being executed in the loop.</p>
<p>Inside the <code>process</code> function, we simply say the process with the given <code>pid</code> has started as a print statement, sleeps for a second, and then print the end of the process with <code>pid</code>. The end of the function is marked with <code>defer wg.Done()</code>, this is to indicate that the go routine is completed.</p>
<p>The wg(wait group) has a counter that keeps track of the number of go routines it has to synchronize or wait for them to complete. In the <a href="https://pkg.go.dev/sync#WaitGroup.Add">Add</a> function, the internal counter in the <code>WaitGroup</code> is incremented by the <code>delta</code> the integer parsed as a parameter. And the <a href="https://pkg.go.dev/sync#WaitGroup.Done">Done</a> function, decrements the internal counter in the <code>WaitGroup</code> which indicates that the go routine is completed.</p>
<p>In the main function, we call the <a href="https://pkg.go.dev/sync#WaitGroup.Wait">wg.Wait</a> which will block until the counter for the <code>WaitGroup</code> is 0 i.e. all go routines have completed their execution. So, we created 10 go routines which ran concurrently but were synchronized with the help of WaitGroups. The WaitGroup is allowed to block the main function till all the go routines are done executing.</p>
<h2>Go Routine with WaitGroup Use Cases</h2>
<p>Go routines can be used for creating asynchronous tasks, and also for creating concurrent tasks. By using wait groups, we can create a way to wait for multiple goroutines to complete. By using go routines and wait groups, we can complete n number of tasks in a time frame of 1 task's completion. However, to create concurrent communication between the other processes, we need <code>channels</code> (which we will explore in the next article).</p>
<p>Here is a simple breakdown of what asynchronous and concurrent tasks might refer to:</p>
<p><strong>Asynchronous</strong> tasks run independently of the main program flow, allowing the main program to continue executing without waiting for the task's completion. For example, not blocking other tasks in the sequential flow of the main function.</p>
<p><strong>Concurrent</strong> tasks run simultaneously and can execute at the same time as other tasks. They make use of multiple threads (goroutines in Go's case) to achieve parallel execution. For example, running multiple tasks parallel will cut the time down for spinning each task after the completion of another task.</p>
<p>Some of the asynchronous tasks that can be done with go routines might include the following:</p>
<ul>
<li>Sending a mail while saving the user to the database.</li>
<li>Fetch and process data from multiple websites(web scraping/crawling).</li>
<li>High-performance message brokers and queuing systems for inter-process communication.</li>
</ul>
<p>One practical way of using go routines with wait groups would be to send mail, we are not going to see the actual implementation of the mail-sending stuff. However, we can try to mimic how the setup might be for sending bulk mail. By creating a wait group and a list of mail addresses to send, a go routine can be created with a function that sends those emails.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
)

func sendMail(address string, wg *sync.WaitGroup) {
	fmt.Println(&quot;Sending mail to&quot;, address)
	defer wg.Done()
    // Actual mail sending, smtp stuff
    // handle errors

    // client, err := smtp.Dial(&quot;smtp.example.com:587&quot;)
    // errr = client.Mail(&quot;sender@example.com&quot;)
    // err = client.Rcpt(address)
    // wc, err := client.Data()
    //_, err = wc.Write([]byte(&quot;This is the email body.&quot;))
    // err = wc.Close()
    // client.Quit()
}

func main() {
	emails := []string{
		&quot;recipient1@example.com&quot;,
		&quot;recipient2@example.com&quot;,
		&quot;xyz@example.com&quot;,
	}
	wg := sync.WaitGroup{}
    wg.Add(len(emails))

	for _, email := range emails {
		go sendMail(email, &amp;wg)
	}
	wg.Wait()
	fmt.Println(&quot;All emails queued for sending&quot;)
	// Do other stuff
}
</code></pre>
<p>In the above example, the <code>emails</code> is a list of email ids to send the mail. We have created a wait group that we can initialize with the total number of go routines probably to be executed. The <code>wg.Add</code> method is parsed with the number of emails to be sent hence that equals the go routines to spawn.</p>
<p>So, we can then in the for loop, iterate over each mail and send the emails with the <code>sendMail</code> function as go routine. The <code>wg.Wait</code> function outside the loop, will make sure the main function is halted before all the go routines complete their execution.</p>
<p>There is one more way to call a function as a go routine without changing its signature since we had to pass <code>wg</code> pointer as the WaitGroup reference to it to acknowledge the completion of the go routine. We could wrap these two operations wiz. calling the function and calling the <code>wg.Done</code> method with an anonymous function.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
)

func sendMail(address string) {
	fmt.Println(&quot;Sending mail to&quot;, address)
}

func main() {
	emails := []string{
		&quot;recipient1@example.com&quot;,
		&quot;recipient2@example.com&quot;,
		&quot;xyz@example.com&quot;,
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
	fmt.Println(&quot;All emails queued for sending&quot;)
	// Do other stuff
}
</code></pre>
<p>This does the exact same thing, but we don't have to change the signature of the function, This keeps the functional logic of the function in its place and handles the concurrency on the go.</p>
<pre><code class="language-go">
for _, email := range emails {
    mail := email
    go func(m string) {
        sendMail(m)
        wg.Done()
    }(mail)
}
</code></pre>
<p>If the above bit scares you, don't worry it's too simple.</p>
<ul>
<li>We are iterating over the email slice using a for loop and creating a Go routine for each email address. The loop variable email represents the email address at the current iteration.</li>
<li>However, to avoid the loop variable capture issue (where all Go routines would share the same email variable), we create a new variable mail and assign the value of email to it. This step ensures that each Go routine captures its own copy of the email address.</li>
<li>We immediately create an anonymous function (a closure) using the go keyword. This anonymous function takes the <code>mail</code> variable as a parameter m and is executed concurrently as a Go routine. Inside the Go routine, we call the <code>sendMail</code> function with the email address <code>m</code>.</li>
<li>After the <code>sendMail</code> call has been executed i.e. email has been sent, we call wg.Done() to notify the WaitGroup that the Go routine has completed its work. This allows the WaitGroup to properly synchronize and wait for all Go routines to finish before the program proceeds beyond wg.Wait() in the main function.</li>
</ul>
<p>This is one way to do it if you want to separate the mail-sending logic from the goroutines/concurrency task. However, this should be handled with care as the variables inside the closure might be shared among all the goroutines instead of having individual variable literals.</p>
<p>To ensure that each goroutine operates on its own copy of the email address, we use the approach of creating a new variable mail and passing it as a parameter to the anonymous function. This way, each goroutine captures its unique email address, avoiding any interference or unintended sharing of data between goroutines.</p>
<h2>Mutual Exclusion Locks</h2>
<p>In the previous examples, we saw how goroutines and wait groups allow us to run multiple tasks concurrently in Go. However, sometimes these concurrent goroutines need to access shared resources like memory, files, network sockets, etc.</p>
<p>When more than one goroutine tries to access a resource at the same time, it can lead to <strong>race conditions</strong> and unpredictable behavior. To handle this, we need a way to ensure only one goroutine can access the resource at a time.</p>
<p>This is where mutual exclusion locks come in. A <strong>mutual exclusion lock</strong>, or <strong>mutex</strong>, provides a mechanism to lock access to a shared resource. It ensures only one goroutine at a time can acquire the lock, blocking other goroutines until the lock is released.</p>
<p>For example, say we have multiple goroutines trying to append data to the same memory buffer(could be file/database/etc.) concurrently:</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;os&quot;
	&quot;sync&quot;
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
	go WriteToFile(&quot;data/f1.txt&quot;, &quot;Hello &quot;, &amp;sharedBuffer, &amp;wg)
	go WriteToFile(&quot;data/f1.txt&quot;, &quot;World! &quot;, &amp;sharedBuffer, &amp;wg)
	wg.Wait()

	fmt.Println(string(sharedBuffer))
}
</code></pre>
<pre><code class="language-bash">$ go run --race no-mutex.go
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
</code></pre>
<p>This is because we are trying to access the same memory address at the same time. This is a race condition and can lead to unpredictable behavior. Try removing the <code>--race</code> flag while running, in this little stupid example, it might not be obvious, but in complex and constrained environments, this can get the application in serious trouble.</p>
<p><strong>NOTE</strong>: We are using <code>go run --race no-mutex.go</code> to check if there are any race conditions in the program. This is the <a href="https://go.dev/doc/articles/race_detector">race detector</a> flag in the run command.</p>
<p>To avoid this race condition, we need to add the mutex locks provided in the <a href="https://pkg.go.dev/sync#Mutex">sync.Mutex</a> type. There are methods like <a href="">Lock</a>, <a href="">Unlock</a>, and <a href="">TryLock</a> which help in locking access of the resource to a single entity at a given time.</p>
<p>When a goroutine calls <code>Lock()</code> on a mutex, it acquires the lock. If the mutex is already locked by another goroutine, the calling goroutine will be blocked (put to sleep) until the lock becomes available. Once the lock is acquired successfully, the goroutine can proceed with its critical section, which is the part of the code that should not be executed concurrently by multiple goroutines.</p>
<p>When a goroutine calls <code>Unlock()</code> on a mutex, it releases the lock. This allows other waiting goroutines to acquire the lock and execute their critical sections. It's essential to ensure that <code>Unlock()</code> is called after the critical section to release the mutex and avoid deadlocks. The critical section/shared resource should not be accessed after the release of this lock.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;os&quot;
	&quot;sync&quot;
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
	go WriteToFile(&quot;data/f1.txt&quot;, &quot;Hello Gophers!
&quot;, &amp;sharedBuffer, &amp;wg, &amp;mut)
	go WriteToFile(&quot;data/f1.txt&quot;, &quot;Welcome to Goworld!
&quot;, &amp;sharedBuffer, &amp;wg, &amp;mut)
	wg.Wait()

	fmt.Println(string(sharedBuffer))
}
</code></pre>
<pre><code class="language-bash">$ go run --race mutex.go
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
</code></pre>
<p>The above example is a preventive measure for race conditions with a mutex Lock on the shared resource used in the go routines.</p>
<p>Let's break down the code step by step. First, we initialize a few variables:</p>
<ul>
<li>
<p><code>wg</code> as <code>sync.WaitGroup</code>:
The <code>wg</code> is a waitgroup that will be used for the synchronization of the go-routines by blocking the main function to exit</p>
</li>
<li>
<p><code>mut</code> as <code>sync.Mutex</code>:
The <code>mut</code> is a structure that internally holds a few integer values indicating either a blocked or unblocked state. The <code>sync.Mutex</code> has two private fields wiz. <code>state</code> and <code>sema</code>, the state holds the mutex state either <code>0</code>(unlocked) or <code>1</code> as locked. The <code>sema</code> field is a <code>uint32</code> that is used for blocking and signaling, it acts as a semaphore to manage blocking and unblocking goroutines trying to acquire the mutex.
This will be used to acquire <code>Lock</code> and <code>Unlock</code> on the shared resource while writing the data to the file or appending the data to the resource.</p>
</li>
<li>
<p><code>sharedBuffer</code> as <code>[]byte</code>:
The <code>sharedBuffer</code> is the actually shared resource that will be used to hold the strings for keeping track of the data written to the file. It will be the resource that would require to lock for mutating its value(appending to the slice) among the go routines.</p>
</li>
</ul>
<p>We add <code>2</code> to the <code>wg</code> to indicate to wait for the two go-routines to complete, in the next line we call two go routines as the function <code>WriteToFile</code>. The <code>WriteToFile</code> is a function that takes in quite a few parameters namely, the filename, the content to write, the reference to the sharedBuffer, waitgroup, and the mutex.</p>
<p>Inside the function <code>WriteToFile</code>:</p>
<ul>
<li>We first <code>defer</code> the <code>waitgroups</code> as <code>Done</code> i.e. to call the <code>wg.Done</code> method at the end of the function call.</li>
<li>Typecast the <code>contents</code> from <code>string</code> as the <code>[]byte</code>.</li>
<li>Acquire the <code>mutex.Lock()</code> i.e. to say &quot;The below operations should not be done concurrently, one at a time&quot;. We then append the <code>contents</code> to the <code>buffer</code> which is a pointer to the <code>sharedBuffer</code> in the main function, so essentially we are trying to mutate the <code>sharedBuffer</code> in this function.</li>
<li>Open the file as <code>O_APPEND</code> and <code>O_WRONLY</code> to indicate that the file should be opened in append/writing mode. (We have observed this type of operation in the <a href="https://www.meetgor.com/golang-file-write/">Golang: File Write</a> article)</li>
<li>Use the <code>Write</code> method to write the slice of bytes(contents) into the file that we opened. We have a <code>defer</code> for the closing of the file.</li>
</ul>
<p>We obviously check for errors and print them, but it could be a <code>panic</code> or <code>log</code> depending on the type of operation the application is doing.
So that is the all operation we want to do, so we finally open the lock with <code>mutex.Unlock()</code> which will allow the other go routine if any to access the resource and proceed with the operations.</p>
<h3>Read Write Mutual Exclusion Lock</h3>
<p>The mutual exclusion lock is good if you have a write operation-heavy application. However, if we have read write operations in more or the less in same proportion(read heavy) we don't want to have readers getting blocked when other readers are accessing the resource since it is not a mutation process.</p>
<p>We could allow many readers to read simultaneously. But for the writing operation, we want to block the readers/writers. The writer should be given preference first in case of a lock by other writers. This would prevent a writer from waiting for readers to complete. This is usually referred to as <strong>Read Write Mutual Exclusion</strong>.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
	&quot;time&quot;
)

func reader(id int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	if !mutex.TryLock() {
		fmt.Printf(&quot;Reader %d blocked!
&quot;, id)
		return
	}
	defer mutex.Unlock()
	fmt.Printf(&quot;Reader %d: read count %d
&quot;, id, *count)
}

func writer(id, increment int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	*count += increment
	time.Sleep(5 * time.Millisecond)
	fmt.Printf(&quot;Writer %d: wrote count %d
&quot;, id, *count)
}

func main() {

	count := 1
	var wg sync.WaitGroup
	var mutex sync.Mutex
    readers := 5
    writers := 3
	wg.Add(readers)
	for i := 0; i &lt; readers; i++ {
		go reader(i, &amp;count, &amp;wg, &amp;mutex)
	}

	wg.Add(writers)
	for i := 0; i &lt; writers; i++ {
		go writer(i, 1, &amp;count, &amp;wg, &amp;mutex)
	}
	wg.Wait()

}
</code></pre>
<pre><code class="language-bash">$ go run --race no-rwmutex.go 
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
</code></pre>
<p>The above example has a <code>reader</code> and a <code>writer</code> method, the <code>reader</code> method simply has to read the shared resource <code>count</code>. It acquires a <code>Mutex</code> lock before reading and unlocks it thereafter. Similarly, the <code>writer</code> function is used for incrementing the <code>count</code> shared resource.</p>
<p>The <code>reader</code> method has a <a href="https://pkg.go.dev/sync#Mutex.TryLock">TryLock</a> method that tries to acquire a mutex lock on the resource, if the resource is already locked, the function will return <code>false</code> and hence we can say that the reading is blocked(just for demonstration). And if the function <code>TryLock</code> returns <code>true</code>, it will acquire the <code>Lock</code>. We further <code>defer</code> the <code>Unlock</code> and access the <code>count</code> variable which is passed as a reference.</p>
<p>The <code>writer</code> method is simply acquiring the <code>Lock</code> and incrementing the <code>counter</code> and thereafter <code>Unlock</code> is called with <code>defer</code>.</p>
<p>In the above code:</p>
<ul>
<li>The <code>reader</code> and the <code>writer</code> both might be waiting for the lock to be released, however, for readers to wait for reading doesn't make sense.</li>
<li>Because if you would want to just read a particular memory address, there shouldn't be any locks for one reader to wait for other readers to finish.</li>
<li>However, for writing, there has to be a lock. The mutex lock will lock the resource irrespective of the <code>reader</code> or <code>writer</code>.</li>
</ul>
<p>This might not be as visible here, but it might be the reason, that all the <code>readers</code> are blocked from reading due to another reader's or writer's lock.</p>
<p>The <code>sync</code> package has the <a href="https://pkg.go.dev/sync#RWMutex">RWMutex</a> that does this exact same thing. It is almost similar to the <code>Mutex</code> however, it would allow concurrent reading operation and prefer writing operation before readers to prevent writer starvation.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
	&quot;time&quot;
)

func reader(id int, count *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	defer mutex.RUnlock()

	if !mutex.TryRLock() {
		fmt.Printf(&quot;Reader %d blocked!
&quot;, id)
		return
	}
	fmt.Printf(&quot;Reader %d: read count %d
&quot;, id, *count)

}

func writer(id, increment int, count *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()

	*count += increment
	time.Sleep(5 * time.Millisecond)
	fmt.Printf(&quot;Writer %d: wrote count %d
&quot;, id, *count)
}

func main() {

	count := 1
	var wg sync.WaitGroup
	var mutex sync.RWMutex
    readers := 5
    writers := 3
	wg.Add(readers)
	for i := 0; i &lt; readers; i++ {
		go reader(i, &amp;count, &amp;wg, &amp;mutex)
	}

	wg.Add(writers)
	for i := 0; i &lt; writers; i++ {
		go writer(i, 1, &amp;count, &amp;wg, &amp;mutex)
	}
	wg.Wait()

}
</code></pre>
<pre><code class="language-bash">$ go run --race rwmutex.go 
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

</code></pre>
<p>In the modified example, all the logic remains the same, just the <code>sync.Mutex</code> is replaced with <code>sync.RWMutex</code>. Also for trying to acquire the lock in the <code>reader</code> method <a href="https://pkg.go.dev/sync#RWMutex.RLock">TryRLock</a> is used instead of <a href="https://pkg.go.dev/sync#RWMutex.TryRLock">TryLock</a> which will check if the existing lock acquired is of a reader or writer, if it is a reader, it will return <code>true</code>, else <code>false</code>. Also the <code>Unlock</code> is replaced with the <code>RUnlock</code> method for releasing the read lock. In the <code>writer</code> method, everything remains the same so the writer has to acquire the lock irrespective of whether the current lock is from the reader/writer, so it is a normal <code>Lock</code> and <code>Unlock</code>.</p>
<p>In the above example, we can see all the read operations sometimes are executed instead of getting blocked. This is due to the <code>RWMutex</code> Lock and Unlock on the read operation/function.</p>
<ul>
<li>When one reader is reading, it can't block other readers.</li>
<li>However with simple <code>Mutex</code>, the reader is even blocked when another reader is reading.</li>
<li>For the write operation, it will be blocked as usual, so if a writer is performing a write operation and a reader/readers come in, they will be blocked, also if in the meantime, while the resource is locked, another writer comes in, the writer will be given preference instead of waiting for all readers to complete. This prevents writer starvation.</li>
</ul>
<p>You can see, that the readers are still blocked the majority of the time, however, they are blocked due to the writer locking the resource and not any other reader.</p>
<p>There is a bit more suitable <a href="https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/go-routines/file-rw.go">example</a> related to <code>file</code> reading and writing with <code>RWMutex</code>, make sure to check that out to get a more clear understanding of an actual case of using the <code>RWMutex</code>.</p>
<h3>Channels</h3>
<p>This is a big part and I would like to delve into this topic in a separate post. There are some patterns like <code>fan-in</code>, <code>fan-out</code>, <code>worker-pool</code>, <code>pub-sub</code>, etc. which are really common in web applications and backend systems. These patterns we shall explore in the next article.</p>
<p>Channels are a way to provide a safe and idiomatic way for Goroutines to exchange data and coordinate their execution without resorting to low-level mechanisms like shared memory or explicit locking.</p>
<p>That's it from the 30th part of the series, all the source code for the examples are linked in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/go-routines">100 days of Golang</a> repository.</p>
<h3>References:</h3>
<ul>
<li><a href="https://blog.nindalf.com/posts/how-goroutines-work/">How Goroutines Work</a></li>
<li><a href="https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/">Concurrency patterns in Golang: WaitGroups and Goroutines</a></li>
<li><a href="https://riteeksrivastava.medium.com/a-complete-journey-with-goroutines-8472630c7f5c">A complete journey with Goroutines</a></li>
</ul>
<h2>Conclusion</h2>
<p>From this part of the series, the fundamentals of golang's concurrency model were understood specifically spawning go-routines, synchronously executing go-routines with the help of a wait group, mutex locks, and how to secure concurrent access to a shared resource. In the next part of the series, these concepts will be used in asynchronous communication using channels.</p>
<p>Hopefully, you have got the basics of concurrency in golang cleared from this post. If you have any queries, suggestions, or feedback, please feel free to comment below or contact me on the social handles. Thank you for reading. Happy Coding :)</p>
