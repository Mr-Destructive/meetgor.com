{
  "title": "Golang: Test Output JSON",
  "status": "published",
  "slug": "golang-test-output-json",
  "date": "2025-04-05T12:33:28Z"
}

<p>I just discovered that we can generate a JSON output of test results in Golang. I found this <a href="https://youtu.be/cf72gMBrsI0?t=80">here</a>.</p>
<p>Let's take a fresh simple example.</p>
<pre><code class="language-go">package jsontest

func hello() string {
    return &quot;Hello, World!&quot;
}
</code></pre>
<pre><code class="language-go">package jsontest

import (
    &quot;testing&quot;
)

func TestHello(t *testing.T) {
    want := &quot;Hello, World!&quot;
    got := hello()
    t.Logf(&quot;got: %q&quot;, got)
    if got != want {
        t.Errorf(&quot;got: %q, want: %q&quot;, got, want)
    }
}
</code></pre>
<p>Here, we have a function <code>hello</code> that simply returns a string, the <code>TestHello</code> function in the <code>jsontest</code> package will check if the returned string is correctly returned or not.</p>
<p>So, we can test this with <code>go test ./...</code> command, this will give out the output in a standard output/error in text format. However, if we add the <code>-json</code> flag, we can get the output in JSON format.</p>
<pre><code class="language-bash">go test ./... -json
</code></pre>
<pre><code class="language-json">{&quot;Time&quot;:&quot;2024-01-01T20:52:45.861974085+05:30&quot;,&quot;Action&quot;:&quot;start&quot;,&quot;Package&quot;:&quot;json-test&quot;}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863133332+05:30&quot;,&quot;Action&quot;:&quot;run&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Test&quot;:&quot;TestHello&quot;}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863142397+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Output&quot;:&quot;=== RUN   TestHello\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863148346+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Output&quot;:&quot;    main_test.go:10: got: \&quot;Hello, World\&quot;\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863151351+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Output&quot;:&quot;    main_test.go:12: got: \&quot;Hello, World\&quot;, want: \&quot;Hello, World!\&quot;\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863157014+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Output&quot;:&quot;--- FAIL: TestHello (0.00s)\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863160418+05:30&quot;,&quot;Action&quot;:&quot;fail&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Elapsed&quot;:0}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863411555+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Output&quot;:&quot;FAIL\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863438344+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Output&quot;:&quot;FAIL\tjson-test\t0.001s\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T20:52:45.863443461+05:30&quot;,&quot;Action&quot;:&quot;fail&quot;,&quot;Package&quot;:&quot;json-test&quot;,&quot;Elapsed&quot;:0.001}
</code></pre>
<p>Pretty cool right?</p>
<p>This is really useful for programmatically taking the output and parsing it to get the metrics.</p>
<p>We can even combine with the coverage flag to get the coverage metrics as well.</p>
<pre><code class="language-bash">go test ./... -json -cover
</code></pre>
<pre><code class="language-json">{&quot;Time&quot;:&quot;2024-01-01T21:13:30.771976961+05:30&quot;,&quot;Action&quot;:&quot;start&quot;,&quot;Package&quot;:&quot;jsontest&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.775118482+05:30&quot;,&quot;Action&quot;:&quot;run&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestHello&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.775172535+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Output&quot;:&quot;=== RUN   TestHello\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.775201647+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Output&quot;:&quot;    main_test.go:10: got: \&quot;Hello, World!\&quot;\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.775231759+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Output&quot;:&quot;--- PASS: TestHello (0.00s)\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.775253928+05:30&quot;,&quot;Action&quot;:&quot;pass&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestHello&quot;,&quot;Elapsed&quot;:0}
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.775269402+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Output&quot;:&quot;PASS\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.776153185+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Output&quot;:&quot;coverage: 100.0% of statements\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.776808599+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Output&quot;:&quot;ok  \tjsontest\t0.004s\tcoverage: 100.0% of statements\n&quot;
{&quot;Time&quot;:&quot;2024-01-01T21:13:30.777814589+05:30&quot;,&quot;Action&quot;:&quot;pass&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Elapsed&quot;:0.006}
</code></pre>
<p>I am planning to use this in my workflow for integrating the output of the tests suite with specific tests.</p>
<p>For running the specific tests, you can use <code>go test -run TestName</code> command, this will only run the provided test function.</p>
<pre><code class="language-go">// main.go

package jsontest

func hello() string {
    return &quot;Hello, World!&quot;
}

func add(x, y int) int {
    return x + y
}
</code></pre>
<pre><code class="language-go">// main_test.go
package jsontest

import (
	&quot;testing&quot;
)

func TestHello(t *testing.T) {
    want := &quot;Hello, World!&quot;
    got := hello()
    t.Logf(&quot;got: %q&quot;, got)
    if got != want {
        t.Errorf(&quot;got: %q, want: %q&quot;, got, want)
    }
}

func TestAdd(t *testing.T) {
    want := 2
    got := add(1, 1)
    t.Logf(&quot;got: %d&quot;, got)
    if got != want {
        t.Errorf(&quot;got: %d, want: %d&quot;, got, want)
    }
}
</code></pre>
<p>So, we have two test in this go module, we can run a specific test using <code>go test -run TestName</code> command as so:</p>
<pre><code class="language-bash">go test -run TestAdd -json
</code></pre>
<pre><code class="language-json">{&quot;Time&quot;:&quot;2024-01-01T21:33:44.19397581+05:30&quot;,&quot;Action&quot;:&quot;start&quot;,&quot;Package&quot;:&quot;jsontest&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:33:44.198067398+05:30&quot;,&quot;Action&quot;:&quot;run&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestAdd&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:33:44.198150156+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestAdd&quot;,&quot;Output&quot;:&quot;=== RUN   TestAdd\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:33:44.198197444+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestAdd&quot;,&quot;Output&quot;:&quot;    main_test.go:19: got: 2\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:33:44.198217057+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestAdd&quot;,&quot;Output&quot;:&quot;--- PASS: TestAdd (0.00s)\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:33:44.198230965+05:30&quot;,&quot;Action&quot;:&quot;pass&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Test&quot;:&quot;TestAdd&quot;,&quot;Elapsed&quot;:0}
{&quot;Time&quot;:&quot;2024-01-01T21:33:44.198241628+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Output&quot;:&quot;PASS\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:33:44.19869148+05:30&quot;,&quot;Action&quot;:&quot;output&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Output&quot;:&quot;ok  \tjsontest\t0.004s\n&quot;}
{&quot;Time&quot;:&quot;2024-01-01T21:33:44.198822637+05:30&quot;,&quot;Action&quot;:&quot;pass&quot;,&quot;Package&quot;:&quot;jsontest&quot;,&quot;Elapsed&quot;:0.005}
</code></pre>
<p>As we can see, there is only one test being executed and the output of the test is in JSON format.</p>
<p>These are really good flags and options to have as they make the output more portable. I will be planning to use this to improve my workflow in testing and developing open source projects and personal projects as well. I am really inspired by the Teej's video of executing anything in NeoVim.</p>
