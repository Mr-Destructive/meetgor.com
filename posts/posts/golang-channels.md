{
  "title": "Golang: Channels",
  "status": "published",
  "slug": "golang-channels",
  "date": "2025-04-05T12:33:30Z"
}

<h2>Introduction</h2>
<p>In this part of the series, we will be continuing with the concurrency features of golang with channels. In the last post, we covered the fundamentals of go routines and wait groups. By leveraging those understood concepts, we will explore channels to communicate the data between various go routines.</p>
<h2>What are Channels</h2>
<p>A golang Channel is like a pipe that lets goroutines communicate. It lets you pass values from one goroutine to another. Channels are typed i.e. you declare them with <code>chan</code> keyword followed by the type to be sent and received (e.g. <code>chan int</code>). The <code>chan</code> type specifies the type of values that will be passed through the channel. We will explore the detailed technicalities soon. Right now, we need to just focus on what problem is channels solving.</p>
<p>In the previous article, we worked with go routines and wait groups which allowed us to process tasks asynchronously. However, if we wanted to access the data in between the processes, we would have to tweak the core functionality or might require global variables, however, in real-world applications, the environment is quite constrained. We would require a way to communicate data between those go routines. Channels are made just for that(more than that), but in essence, it solves that exact problem.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

func main() {
	ch := make(chan string)
	defer close(ch)

	go func() {
		message := &quot;Hello, Gophers!&quot;
		ch &lt;- message
	}()

	msg := &lt;-ch
	fmt.Println(msg)
}
</code></pre>
<p>In the above code example, the channel <code>ch</code> is created of type <code>string</code> and a message is sent to the channel inside a go routine as <code>ch &lt;- message</code>, and the message is retrieved from the channel as <code>&lt;-ch</code>.</p>
<pre><code class="language-bash">$ go run main.go

Hello, Gophers!
</code></pre>
<p>Channels have two key properties:</p>
<ul>
<li>Send and receive operations block until both sides are ready(i.e. there is a sender and a receiver for a channel). This allows goroutines to synchronize without explicit locks or condition variables.</li>
<li>Channels are typed, so only values of the specified type can be sent and received. This provides type safety.</li>
</ul>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
)

func main() {
    ch := make(chan string)
    go func() {
        message := &quot;Hello, Gophers!&quot;
        ch &lt;- message
    }()
    fmt.Println(&lt;-ch)
    fmt.Println(&lt;-ch)
}
</code></pre>
<p>In the same example, if we tried to add the second receiver i.e. <code>&lt;-ch</code>, it would result in a deadlock/block forever since there is no second message sent into the channel. Only one value i.e. &quot;Hello Gophers!&quot; was sent as a <code>message</code> into the channel, and that was received by the first receiver as <code>&lt;-ch</code>, however in the next receiver, there is no sender.</p>
<pre><code class="language-bash">$ go run main.go

Hello, Gophers!
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /home/meet/code/100-days-of-golang/scripts/channels/main.go:16 +0x125
exit status 2
</code></pre>
<p>To sum up the deadlock concept in unbuffered channels:</p>
<ul>
<li>The main goroutine is waiting at the second receive operation for a second message that will never arrive (was never sent).</li>
<li>The anonymous goroutine is waiting for someone to read from the channel so it can proceed with sending the second message.</li>
</ul>
<h2>Buffered Channels</h2>
<p>In Go, you can create both buffered and unbuffered channels. An unbuffered channel has no capacity to hold data, it relies on immediate communication between the sender and receiver. However, you can create a buffered channel by specifying a capacity when using the make function, like <code>ch := make(chan int, 5)</code> will create a channel with a capacity of <code>5</code> i.e. it can store a certain number of values without an immediate receiver. A buffered channel allows you to send multiple values to the channel without an immediate receiver, up to its capacity. After that, it will block until the receiver retrieves values.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
)

func main() {
	buffchan := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 1; i &lt;= 2; i++ {
		go func(n int) {
			buffchan &lt;- n
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(buffchan)

	for c := range buffchan {
		fmt.Println(c)
	}
}
</code></pre>
<pre><code class="language-bash">$ go run channels.go
1
2

$ go run channels.go
2
1
</code></pre>
<p>In this code snippet, we create a buffered channel ch with a capacity of 2. We send two values to the channel, and even though there's no immediate receiver, the code doesn't block. If we were to send a third value, it would lead to a deadlock because there is no receiver to free up space in the buffer.</p>
<h2>Closing Channels</h2>
<p>Closing a channel is important to signal to the receiver that no more data will be sent. It's achieved using the built-in close function. After closing a channel, any further attempts to send data will result in a panic. On the receiving side, if a channel is closed and there's no more data to receive, the receive operation will yield the zero value for the channel's type.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i &lt;= 5; i++ {
			ch &lt;- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(&quot;Received:&quot;, num)
	}
}
</code></pre>
<p>In this example, a goroutine sends numbers to the channel and then closes it. The main routine receives these numbers using a for-range loop. When the channel is closed and all values are received, the loop will terminate automatically. Keep in mind that only a sender can close the channel, to indicate the receiver to not wait for further values from the channel.</p>
<h2>Select Statement for Channels</h2>
<p>The select statement is used for handling multiple channels. There are a few operations that can be checked with a case statement in the select block.</p>
<p>|Case     |Channel Operation|
|---------|-----------------|
|         |                 |
|Sending  | chan &lt;- value   |
|         |                 |
|Receiving|    &lt;- chan      |
|         |                 |</p>
<p>So, we can either check if there is a sender or a receiver available for a channel with a case statement just like a switch statement.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
)

func sendMessage(ch chan string, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch &lt;- message
}

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)
	wg.Add(2)

	go sendMessage(ch1, &quot;Hello, Gophers!&quot;, &amp;wg)
	go sendMessage(ch2, &quot;Hello, Hamsters!&quot;, &amp;wg)

	go func() {
		defer wg.Done()
		wg.Wait()
		close(ch1)
		close(ch2)
	}()
	ch1 &lt;- &quot;new message to c1&quot;
	ch2 &lt;- &quot;new message to c2&quot;

	select {
	case &lt;-ch1:
		fmt.Println(&quot;Received from ch1&quot;)
	case ch1 &lt;- &quot;new message to c1&quot;:
		fmt.Println(&quot;Sent to ch1&quot;)
	case &lt;-ch2:
		fmt.Println(&quot;Received from ch2&quot;)
	case ch2 &lt;- &quot;new message to c2&quot;:
		fmt.Println(&quot;Sent to ch2&quot;)
	}
}
</code></pre>
<pre><code class="language-bash">$ go run channels.go
Sent to ch1

$ go run simple.go
Received from ch1

$ go run simple.go
Received from ch2

$ go run simple.go
Sent to ch2

$ go run simple.go
Received from ch1
</code></pre>
<p>The order of the messages is not guaranteed, the operation which is performed first based on the go routine will be only logged.</p>
<p>In the simple example above, we have created two channels <code>ch1</code> and <code>ch2</code>, and sent two messages to them using two go routines. The main routine then waits for the messages to be received from the channels. We close the channels when the sending is done and simply check for the 4 cases i.e. the send on channel 1, receive on channel 1, and similarly for channel 2. So, that is how we can use the select statement to check which operation is being performed on the different channels, and this forms the basis for the communication between channels. We get more ease in the flow while working with channels.</p>
<p>Below is an example to test which <code>url</code> or a web server is responding first to the request.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
	&quot;sync&quot;
)

func pingGoogle(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get(&quot;http://google.com&quot;)
	c &lt;- res.Status
}

func pingDuckDuckGo(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get(&quot;https://duckduckgo.com&quot;)
	c &lt;- res.Status
}

func pingBraveSearch(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get(&quot;https://search.brave.com&quot;)
	c &lt;- res.Status
}

func main() {
	gogChan := make(chan string)
	ddgChan := make(chan string)
	braveChan := make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)

	go pingDuckDuckGo(ddgChan, &amp;wg)
	go pingGoogle(gogChan, &amp;wg)
	go pingBraveSearch(braveChan, &amp;wg)

	openChannels := 3

	go func() {
		wg.Wait()
		close(gogChan)
		close(ddgChan)
		close(braveChan)
	}()

	for openChannels &gt; 0 {
		select {
		case msg1, ok := &lt;-gogChan:
			if !ok {
				openChannels--
			} else {
				fmt.Println(&quot;Google responded:&quot;, msg1)
			}
		case msg2, ok := &lt;-ddgChan:
			if !ok {
				openChannels--
			} else {
				fmt.Println(&quot;DuckDuckGo responded:&quot;, msg2)
			}
		case msg3, ok := &lt;-braveChan:
			if !ok {
				openChannels--
			} else {
				fmt.Println(&quot;BraveSearch responded:&quot;, msg3)
			}
		}
	}
}
</code></pre>
<p>The above example shows how to use a select statement to wait for multiple channels to be ready before proceeding with the next operation. With this example, we can get the channel that sent the response first i.e. which search engine in this case responded to the ping first. Just a bit exaggerated example but it helps in understanding the concept of the <code>select</code> statement.</p>
<pre><code class="language-bash">$ go run select-chan.go

DuckDuckGo responded: 200 OK
Google responded: 200 OK
BraveSearch responded: 200 OK


$ go run select-chan.go

DuckDuckGo responded: 200 OK
BraveSearch responded: 200 OK
Google responded: 200 OK
</code></pre>
<p>Let's break each of the steps down:</p>
<ul>
<li><code>pingDuckDuckGo(ddgChan, &amp;wg)</code> is a method which sends data to the channel <code>ddgChan</code>.</li>
<li><code>pingGoogle(gogChan, &amp;wg)</code> is a method that sends data to the channel <code>gogChan</code>.</li>
<li><code>pingBraveSearch(braveChan, &amp;wg)</code> is a method that sends data to the channel <code>braveChan</code>.</li>
<li>We wait for each go routine to finish using <code>wg.Wait()</code> and close the channels.</li>
<li>Finally, we close the channels <code>gogChan</code>, <code>ddgChan</code>, and <code>braveChan</code> to pick up the data from the channel as <code>&lt;-chan</code> with the select case block.</li>
<li>The select case will pick the first channel that is ready to receive data. Hence we get the output based on the order of which the channel responded first.</li>
<li>We use the <code>!ok</code> condition to check if the channel is closed or not, we have a <code>openChannels</code> variable to keep track of the number of open channels, if there are no channels open, we simply break out of the infinite loop.</li>
</ul>
<h2>Directional Channels</h2>
<p>Channels can also be designated as &quot;send-only&quot; or &quot;receive-only&quot; to enforce certain communication patterns and enhance safety. This is done by specifying the direction when defining the channel type.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
)

func receiver(ch &lt;-chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println(&quot;Received:&quot;, i)
	}
	wg.Done()
}

func sender(ch chan&lt;- int, wg *sync.WaitGroup) {
	for i := 0; i &lt; 10; i++ {
		fmt.Println(&quot;Sent:&quot;, i)
		ch &lt;- i
	}
	close(ch)
	wg.Done()
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go receiver(ch, &amp;wg)
	go sender(ch, &amp;wg)
	wg.Wait()
}
</code></pre>
<p>In the above example, we have created a channel <code>ch</code> and sent 10 values to it using two-go routines. The main routine is waiting for the goroutines to finish before closing the channel. The <code>sender</code> sends values <code>0</code> through <code>9</code>, and the <code>receiver</code> prints whenever a value is received. In the <code>sender</code> method, we only accept the channel to send data as <code>chan&lt;-</code>, and in the <code>receiver</code> method, the channel parameter is set to only read from the channel as <code>&lt;-chan</code>.</p>
<pre><code class="language-bash">$ go run send-recv.go

Sent: 0
Received: 0
Sent: 1
Sent: 2
Received: 1
Received: 2
Sent: 3
Sent: 4
Received: 3
Received: 4
Sent: 5
Sent: 6
Received: 5
Received: 6
Sent: 7
Sent: 8
Received: 7
Received: 8
Sent: 9
Received: 9
</code></pre>
<p>When we define a parameter as a write-only channel means that the function can only send data into that channel. It cannot read data from it or close it. This pattern is helpful when you want to make sure that the function is solely responsible for producing data and not consuming or interacting with the channel's current state.</p>
<p>When we define a parameter as a read-only channel, it means that the function can only receive data from that channel. It cannot close the channel or send data into it. This pattern is useful when we want to ensure that the function only consumes data from the channel without modifying it or interfering with the sender's logic.</p>
<p>Additionally, the compiler will catch code trying to send on a read-only channel or receive from a write-only one.</p>
<h2>Common Channel Usage Pattern</h2>
<p>There are a variety of ways in which channels can be used in Go. In this section, we'll explore some of the most common patterns for using channels in Go. Some of the most useful and idiomatic channel usage patterns include pipelines, fan-in and fan-out, etc.</p>
<h3>Async Await pattern for Channels</h3>
<p>In Go, goroutines and channels enable an elegant async/await style. A goroutine can execute a task asynchronously, while the main thread awaits the result using a channel.</p>
<p>The async-await pattern in Go involves initiating multiple tasks concurrently, each with its own goroutine, and then awaiting their completion before proceeding. Channels are used to communicate between these goroutines, allowing them to work independently and provide results to the main routine when ready.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func fetchURL(url string, ch chan&lt;- http.Response) {
	go func() {
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		ch &lt;- *res
	}()
}

func task(name string) {
	fmt.Println(&quot;Task&quot;, name)
}

func main() {
	fmt.Println(&quot;Start&quot;)

	url := &quot;http://google.com&quot;

	respCh := make(chan http.Response)

	fetchURL(url, respCh)

	task(&quot;A&quot;)
	task(&quot;B&quot;)

	response := &lt;-respCh
	fmt.Println(&quot;Response Status:&quot;, response.Status)

	fmt.Println(&quot;Done&quot;)
}
</code></pre>
<pre><code class="language-bash">$ go run async.go
Start
Task A
Task B
Response Status: 200 OK
Done
</code></pre>
<p>In the above example, we have created a function <code>fetchURL</code> which takes a URL and a channel as an argument. The channel <code>respCh</code> is used to communicate between the goroutines. The function fires up a goroutine that fetches the request, we send a <code>GET</code> request to the provided URL and send the response to the provided channel.  In the main function, we access the <code>response</code> by receiving the data from the channel as <code>&lt;-respCh</code>. Before doing this, we can do any other task simultaneously, like <code>task(&quot;A&quot;)</code> and <code>task(&quot;B&quot;)</code> which just prints some string(it could be anything). But this should be before we pull in from the channel, anything after the access will be blocked i.e. will be executed sequentially.</p>
<h3>Pipeline pattern for Channels</h3>
<p>The pipeline pattern is used to chain together a sequence of processing stages, each stage consumes input, processes data, and passes the output to the next stage. This type of pattern can be achieved by chaining different channels from one go routine to another.</p>
<p><img src="https://meetgor-cdn.pages.dev/100-days-of-golang/channels-pipelines-pattern.png" alt="Pipeline pattern flow using channels in golang"></p>
<p>So, the pipeline pattern using channels in Go, data flows sequentially through multiple stages: Stage 1 reads input and sends to Channel A, Stage 2 receives from Channel A and sends to Channel B, and Stage 3 receives from Channel B and produces the final output.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;sync&quot;
)

func generate(nums []int, out chan&lt;- int, wg *sync.WaitGroup) {
	fmt.Println(&quot;Stage 1&quot;)
	for _, n := range nums {
		fmt.Println(&quot;Number:&quot;, n)
		out &lt;- n
	}
	close(out)
	wg.Done()
}

func square(in &lt;-chan int, out chan&lt;- int, wg *sync.WaitGroup) {
	fmt.Println(&quot;Stage 2&quot;)
	for n := range in {
		sq := n * n
		fmt.Println(&quot;Square:&quot;, sq)
		out &lt;- sq
	}
	close(out)
	wg.Done()
}

func print(in &lt;-chan int, wg *sync.WaitGroup) {
	for n := range in {
		fmt.Println(n)
	}
	wg.Done()
}

func main() {
	input := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	wg.Add(3)

	stage1 := make(chan int)
	stage2 := make(chan int)

	go generate(input, stage1, &amp;wg)

	go square(stage1, stage2, &amp;wg)

	go print(stage2, &amp;wg)

	wg.Wait()
}
</code></pre>
<p>In the above example, we have created a sequence of processing stages, each stage consumes input, processes data, and passes the output to the next stage. We can consider the functions <code>generate</code>, <code>square</code>, and <code>print</code> as stages <code>1</code>, <code>2</code>, and <code>3</code> respectively.</p>
<ul>
<li>The generate function, takes in the input as a slice of integers, an unbuffered channel, and the waitgroup b reference, the function basically iterates over the numbers in the slice and sends it to the channel provided in the parameters.</li>
<li>The square function takes in the stage1 channel that the channel from the stage1, as well as its own channel as <code>stage2</code> (remember the <code>stage1</code> channel has sent the numbers via the generating function).</li>
<li>The square function then iterates over the number sent from the channel <code>stage1</code> as <code>in</code> and squares it and sends it to the channel provided as <code>stage2</code> as the <code>out</code> channel.</li>
<li>The print function takes in the stage2 channel as an argument and iterates over the number sent from the channel <code>stage2</code> and prints it.</li>
</ul>
<pre><code>$ go run pipeline.go
Stage 1
Number: 1
Stage 2
Square: 1
1
Number: 2
Number: 3
Square: 4
Square: 9
Number: 4
4
9
Square: 16
16
Number: 5
Square: 25
25
</code></pre>
<p>So, we can see the order of the execution, both the pipelines started synchronously, However, they perform the operation only when the data is sent from the previous channel. We first print the <code>number</code> from the <code>generate</code> function, then print the squared value in the <code>square</code> function, and finally print it as <code>Square: value</code> in the print function.</p>
<h3>Fan-In pattern for Channels</h3>
<p>The Fan-In pattern is used for combining data from multiple sources into a single stream for unified processing, often using a shared data structure to aggregate the data. We can create the fan-in pattern by merging multiple input channels into a single output channel.</p>
<p><img src="https://meetgor-cdn.pages.dev/100-days-of-golang/channels-fan-in-pattern.png" alt="Fan-in pattern flow using channels in golang"></p>
<p>The fan-in pattern is when multiple input channels (A, B, C) are read concurrently, and their data is merged into a single output channel (M).</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;io/ioutil&quot;
	&quot;sync&quot;
)

func readFile(file string, ch chan&lt;- string) {
	content, _ := ioutil.ReadFile(file)
	fmt.Println(&quot;Reading from&quot;, file, &quot;as :: &quot;, string(content))
	ch &lt;- string(content)
	close(ch)
}

func merge(chs ...&lt;-chan string) string {
	var wg sync.WaitGroup
	out := &quot;&quot;

	for _, ch := range chs {
		wg.Add(1)
		go func(c &lt;-chan string) {
			for s := range c {
				out += s
			}
			wg.Done()
		}(ch)
	}

	wg.Wait()
	return out
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go readFile(&quot;data/f1.txt&quot;, ch1)
	go readFile(&quot;data/f2.txt&quot;, ch2)

	merged := merge(ch1, ch2)

	fmt.Println(merged)
}
</code></pre>
<p>In the above example, the <code>readFile</code> function reads the contents of the file and sends it to the channels <code>ch1</code> and <code>ch2</code> from different go routines. The <code>readFile</code> takes in the channel as send only channel which reads the file and sends the content to the channel as <code>ch &lt;- string(content)</code>. The <code>merge</code> function takes in <code>2</code> it can also be <code>n</code> number of channels to parse from as indicated as <code>...&lt;-chan</code>, it iterates over each channel, and for each channel, it reads the contents, and appends as a single string.</p>
<pre><code class="language-bash">$ go run fan-in.go

Reading from data/f1.txt as ::  This is from file 1
Reading from data/f2.txt as ::  This is from file 2

This is from file 1
This is from file 2


$ go run fan-in.go
Reading from data/f2.txt as ::  This is from file 2
Reading from data/f1.txt as ::  This is from file 1

This is from file 2
This is from file 1
</code></pre>
<p>So, this is how the fan-in pattern works, We create multiple channels and combine the results into a single stream of data(in this example a single string).</p>
<h3>Fan-Out Pattern for Channels</h3>
<p>The Fan-Out pattern involves taking data from a single source and distributing it to multiple workers or processing units for parallel or concurrent handling. Fan-out design splits an input channel into multiple output channels, it is used to distribute branches of work or data across concurrent processes.</p>
<p><img src="https://meetgor-cdn.pages.dev/100-days-of-golang/channels-fan-out-pattern.png" alt="Fan-Out pattern flow using channels in golang"></p>
<p>The fan-out pattern is when data from a single input channel (A) is distributed to multiple worker channels (X, Y, Z) for parallel processing.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;os&quot;
	&quot;sync&quot;
)

func readFile(file string, ch chan&lt;- string, wg *sync.WaitGroup) {
	defer wg.Done()

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf(&quot;Error reading from %s: %v
&quot;, file, err)
		return
	}

	ch &lt;- string(content)
}

func main() {
	files := []string{&quot;data/f1.txt&quot;, &quot;data/f2.txt&quot;}

	var wg sync.WaitGroup
	ch := make(chan string)

	for _, f := range files {
		wg.Add(1)
		go readFile(f, ch, &amp;wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var fileData []string
	for content := range ch {
		fileData = append(fileData, content)
	}

	fmt.Printf(&quot;Read %d files
&quot;, len(fileData))
	fmt.Printf(&quot;Contents:
%s&quot;, fileData)
}
</code></pre>
<p>In the above example, we create a single channel <code>ch</code> as our single source, we loop over all the files, and create go routines calling the <code>readFile</code> function. The <code>readFile</code> function takes in the filename, channel, and the WaitGroup reference, the function reads the file and sends the content to the channel as <code>ch &lt;- content</code>. The <code>readFile</code> is called concurrently for all the files, Here we have done a fan-out of the task into multiple go routines, then in the main function, we iterate over the channel and receive the content.</p>
<pre><code class="language-bash">$ go run fan-out.go

Read 2 files
Contents:
[This is from file 2
 This is from file 1
]


$ go run fan-out.go

Read 2 files
Contents:
[This is from file 1
 This is from file 2
]
</code></pre>
<p>Here's a brief summary of the fan-out pattern from the example provided:</p>
<ul>
<li>Multiple files are read concurrently using goroutines. This &quot;fans out&quot; the work.</li>
<li>The <code>readFile</code> function runs in a goroutine to process each file separately.</li>
<li>WaitGroup coordinates the goroutines.</li>
<li>A shared channel ch collects the results from each goroutine.</li>
<li>The main goroutine reads from the channel and aggregates the results.</li>
<li>Channel is closed and ranged over to collect results cleanly.</li>
</ul>
<p>I have a few more patterns to demonstrate that have been provided in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/channels/patterns/">100 days of Golang</a> repository.</p>
<p>That's it from the 31st part of the series, all the source code for the examples are linked in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/channels">100 days of Golang</a> repository.</p>
<h2>References</h2>
<ul>
<li><a href="https://go.dev/ref/spec#Channel_types">Channels</a></li>
<li><a href="https://go.dev/doc/effective_go#channels">Effective Go: Channels</a></li>
<li><a href="https://mariocarrion.com/2021/08/19/learning-golang-concurrency-patterns-fan-in-fan-out.html">Fan-In and Fan-Out</a></li>
<li><a href="https://go101.org/article/channel.html">Go 101: Channels</a></li>
</ul>
<h2>Conclusion</h2>
<p>So, from this part of the series, we were able to understand the fundamentals of channels in golang. By using the core concepts from the previous posts like go routines and wait groups, we were able to work with channels in golang. We wrote a few examples for different patterns using concurrency concepts with channels. Patterns like pipelines, fan-in, fan-out, async, and some usage of select statements for channels were explored in this section.</p>
<p>Hopefully, you found this section helpful, if you have any comments or feedback, please let me know in the comments section or on my social handles. Thank you for reading. Happy Coding :)</p>
