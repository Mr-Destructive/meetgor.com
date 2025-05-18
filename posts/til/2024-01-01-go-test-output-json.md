---
type: til
title: "Golang: Test Output JSON"
description: "Obtain JSON output of test results in Golang"
status: published
slug: golang-test-output-json
tags: ["go",]
date: 2024-01-01 21:30:00
---

I just discovered that we can generate a JSON output of test results in Golang. I found this [here](https://youtu.be/cf72gMBrsI0?t=80).

Let's take a fresh simple example.

```go
package jsontest

func hello() string {
    return "Hello, World!"
}
```
```go
package jsontest

import (
    "testing"
)

func TestHello(t *testing.T) {
    want := "Hello, World!"
    got := hello()
    t.Logf("got: %q", got)
    if got != want {
        t.Errorf("got: %q, want: %q", got, want)
    }
}
```

Here, we have a function `hello` that simply returns a string, the `TestHello` function in the `jsontest` package will check if the returned string is correctly returned or not.

So, we can test this with `go test ./...` command, this will give out the output in a standard output/error in text format. However, if we add the `-json` flag, we can get the output in JSON format.


```bash
go test ./... -json
```

```json
{"Time":"2024-01-01T20:52:45.861974085+05:30","Action":"start","Package":"json-test"}
{"Time":"2024-01-01T20:52:45.863133332+05:30","Action":"run","Package":"json-test","Test":"TestHello"}
{"Time":"2024-01-01T20:52:45.863142397+05:30","Action":"output","Package":"json-test","Test":"TestHello","Output":"=== RUN   TestHello\n"}
{"Time":"2024-01-01T20:52:45.863148346+05:30","Action":"output","Package":"json-test","Test":"TestHello","Output":"    main_test.go:10: got: \"Hello, World\"\n"}
{"Time":"2024-01-01T20:52:45.863151351+05:30","Action":"output","Package":"json-test","Test":"TestHello","Output":"    main_test.go:12: got: \"Hello, World\", want: \"Hello, World!\"\n"}
{"Time":"2024-01-01T20:52:45.863157014+05:30","Action":"output","Package":"json-test","Test":"TestHello","Output":"--- FAIL: TestHello (0.00s)\n"}
{"Time":"2024-01-01T20:52:45.863160418+05:30","Action":"fail","Package":"json-test","Test":"TestHello","Elapsed":0}
{"Time":"2024-01-01T20:52:45.863411555+05:30","Action":"output","Package":"json-test","Output":"FAIL\n"}
{"Time":"2024-01-01T20:52:45.863438344+05:30","Action":"output","Package":"json-test","Output":"FAIL\tjson-test\t0.001s\n"}
{"Time":"2024-01-01T20:52:45.863443461+05:30","Action":"fail","Package":"json-test","Elapsed":0.001}
```

Pretty cool right?

This is really useful for programmatically taking the output and parsing it to get the metrics.


We can even combine with the coverage flag to get the coverage metrics as well.

```bash
go test ./... -json -cover
```

```json
{"Time":"2024-01-01T21:13:30.771976961+05:30","Action":"start","Package":"jsontest"}
{"Time":"2024-01-01T21:13:30.775118482+05:30","Action":"run","Package":"jsontest","Test":"TestHello"}
{"Time":"2024-01-01T21:13:30.775172535+05:30","Action":"output","Package":"jsontest","Test":"TestHello","Output":"=== RUN   TestHello\n"}
{"Time":"2024-01-01T21:13:30.775201647+05:30","Action":"output","Package":"jsontest","Test":"TestHello","Output":"    main_test.go:10: got: \"Hello, World!\"\n"}
{"Time":"2024-01-01T21:13:30.775231759+05:30","Action":"output","Package":"jsontest","Test":"TestHello","Output":"--- PASS: TestHello (0.00s)\n"}
{"Time":"2024-01-01T21:13:30.775253928+05:30","Action":"pass","Package":"jsontest","Test":"TestHello","Elapsed":0}
{"Time":"2024-01-01T21:13:30.775269402+05:30","Action":"output","Package":"jsontest","Output":"PASS\n"}
{"Time":"2024-01-01T21:13:30.776153185+05:30","Action":"output","Package":"jsontest","Output":"coverage: 100.0% of statements\n"}
{"Time":"2024-01-01T21:13:30.776808599+05:30","Action":"output","Package":"jsontest","Output":"ok  \tjsontest\t0.004s\tcoverage: 100.0% of statements\n"
{"Time":"2024-01-01T21:13:30.777814589+05:30","Action":"pass","Package":"jsontest","Elapsed":0.006}
```

I am planning to use this in my workflow for integrating the output of the tests suite with specific tests.

For running the specific tests, you can use `go test -run TestName` command, this will only run the provided test function.

```go
// main.go

package jsontest

func hello() string {
    return "Hello, World!"
}

func add(x, y int) int {
    return x + y
}
```

```go
// main_test.go
package jsontest

import (
	"testing"
)

func TestHello(t *testing.T) {
    want := "Hello, World!"
    got := hello()
    t.Logf("got: %q", got)
    if got != want {
        t.Errorf("got: %q, want: %q", got, want)
    }
}

func TestAdd(t *testing.T) {
    want := 2
    got := add(1, 1)
    t.Logf("got: %d", got)
    if got != want {
        t.Errorf("got: %d, want: %d", got, want)
    }
}
```

So, we have two test in this go module, we can run a specific test using `go test -run TestName` command as so:

```bash
go test -run TestAdd -json
```

```json
{"Time":"2024-01-01T21:33:44.19397581+05:30","Action":"start","Package":"jsontest"}
{"Time":"2024-01-01T21:33:44.198067398+05:30","Action":"run","Package":"jsontest","Test":"TestAdd"}
{"Time":"2024-01-01T21:33:44.198150156+05:30","Action":"output","Package":"jsontest","Test":"TestAdd","Output":"=== RUN   TestAdd\n"}
{"Time":"2024-01-01T21:33:44.198197444+05:30","Action":"output","Package":"jsontest","Test":"TestAdd","Output":"    main_test.go:19: got: 2\n"}
{"Time":"2024-01-01T21:33:44.198217057+05:30","Action":"output","Package":"jsontest","Test":"TestAdd","Output":"--- PASS: TestAdd (0.00s)\n"}
{"Time":"2024-01-01T21:33:44.198230965+05:30","Action":"pass","Package":"jsontest","Test":"TestAdd","Elapsed":0}
{"Time":"2024-01-01T21:33:44.198241628+05:30","Action":"output","Package":"jsontest","Output":"PASS\n"}
{"Time":"2024-01-01T21:33:44.19869148+05:30","Action":"output","Package":"jsontest","Output":"ok  \tjsontest\t0.004s\n"}
{"Time":"2024-01-01T21:33:44.198822637+05:30","Action":"pass","Package":"jsontest","Elapsed":0.005}
```

As we can see, there is only one test being executed and the output of the test is in JSON format.

These are really good flags and options to have as they make the output more portable. I will be planning to use this to improve my workflow in testing and developing open source projects and personal projects as well. I am really inspired by the Teej's video of executing anything in NeoVim.

