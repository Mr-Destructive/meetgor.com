{
  "title": "Advent of Code Day 3 in Golang: Do Or Don’t Regex",
  "status": "published",
  "slug": "aoc-2024-day-3",
  "date": "2025-04-05T12:33:19Z"
}

<h2>Introduction</h2>
<p>Well, it is day 3 of the advent of code 2024, and I have been doing it on live streams. I am behind two days but working through them one by one. So far, I have learned a lot of things in Go. Let’s dive in for the day 3.</p>
<h2>Part 1</h2>
<p>Part one to any AOC problem seems straightforward, but as soon as part two is revealed, the real implementation starts to break the sweat (if you weren’t optimistic or thoughtful)</p>
<p>For part 1 for this day, was to parse a string that contained <code>mul(X,Y</code>) an expression, where X could be any 3-digit number. So, there could be multiple such expressions within the string and the purpose was to multiply the X and Y in the individual expression and add them up.</p>
<p><img src="https://meetgor-cdn.pages.dev/aoc-2024-d3-solution-part1.jpg" alt="AOC Day 3 Part 1"></p>
<pre><code class="language-plaintext">
xmul(2,4)&amp;mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
</code></pre>
<p>In this example above there are 4 such expressions, and if we add the multiplications of those, we get the answer as 161.</p>
<h3>Approach</h3>
<p>It looks like a Regex pattern, finding an expression-like pattern in a string. So, the approach would be to find such expressions with a regex pattern and parse the numbers to integers, and multiply them, simply.</p>
<p>You could go ahead and write the parser for iterating over each character in the string and parsing the tokens, then evaluating the expression. That is a valid approach, but I choose to do this because I don’t know how to write a parser honestly, I want to try that solution at the end as well.</p>
<p>But for the first part, a quick regular expression seems a good idea.</p>
<h3>Constructing the Regular Expression</h3>
<p>The first thing is to write the regular expression for the <code>mul(X,Y)</code> part which is the only challenging section in part one. The rest is just simple math.</p>
<p>So, we need to find <code>mul</code>, then a <code>(</code> then any number which is 1 to 3 digits long, then <code>,</code> and again a number that is 1 to 3 digits long, and finally ends with a <code>)</code></p>
<p>That translates to:</p>
<pre><code class="language-plaintext">mul\((\d{1,3}),(\d{1,3})\) 
</code></pre>
<p>Let’s breakdown:</p>
<ul>
<li>
<p><code>mul</code> for capturing the literal word <code>mul</code></p>
</li>
<li>
<p><code>\(</code> this is for the first parenthesis in the expression <code>mul()</code> , we need to escape the bracket in a regular expression if we want to match it, so we use <code>\</code> before it.</p>
</li>
<li>
<p>Then we have a match group <code>(\d{1,3})</code> , this will be the <code>X</code> in <code>mul(X,Y)</code>:</p>
<ul>
<li>
<p>A match group is like a group of a match in a regex, basically, if you want to capture specific parts in the entire match, then we use <code>()</code> to group them individually, this is not necessary but helps in getting the right things without overhead</p>
</li>
<li>
<p>In this case, we are using a match group for capturing a number which can have 1 to 3 digits.</p>
</li>
<li>
<p>The other way to write this is <code>([0-9]{1,3})</code> , which also would do the same thing, (NOTE: there are some differences in <code>[0-9]</code> and <code>\d</code>, but that is very subtle and won’t affect this puzzle, if you are still curious, I prefer reading this <a href="https://unix.stackexchange.com/questions/414226/difference-between-0-9-digit-and-d">StackOverflow question</a>)</p>
</li>
</ul>
</li>
<li>
<p>We then use <code>,</code> for the separator of <code>X</code> and <code>Y</code> operands in the <code>mul(X,Y)</code> expression</p>
</li>
<li>
<p>We then similarly do the match for <code>Y</code> in <code>mul(X,Y)</code> with the <code>(\d{1,3})</code> match group</p>
</li>
<li>
<p>Finally, we end the regular expression with the <code>)</code> to end the expression</p>
</li>
</ul>
<h3>Code it</h3>
<p>This is quite straightforward, we grab the line as a string and use <a href="https://pkg.go.dev/regexp#MustCompile">regexp.MustCompile</a> function which gives us a <a href="https://pkg.go.dev/regexp#Regexp">Regexp</a> object, that in turn has a few methods associated with it to find, match, replace, and other things that can be done with a regular expression on a string.</p>
<p>Once we have the <code>mulRegex</code> , we can use the <a href="https://pkg.go.dev/regexp#Regexp.FindAllStringSubmatch">FindAllStringSubmatch</a> method associated with the <code>Regexp</code> struct in the <code>regexp</code> package. The function takes in a string to perform the regex on, and the maximum number of matches to return. We want all the results, so we pass them in <code>-1</code> to get all the matches.</p>
<p>Now, this method returns a slice of a slice of strings, each slice is a match, and within each slice, there is a slice of string, with the matched string and the subsequent match groups in the string if any.</p>
<pre><code class="language-go">func FindMulExpression(line string) [][]string {
  mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
  return mulRegex.FindAllStringSubmatch(line, len(line))
}
</code></pre>
<p>So, the above function will return something like this</p>
<pre><code class="language-go">[
    [mul(2,4) 2 4]
    [mul(5,5) 5 5]
    [mul(11,8) 11 8]
    [mul(8,5) 8 5]
]
</code></pre>
<p>This is a list of list of strings, these look like numbers but those are string types in Golang.</p>
<p>Now we have this, we can create the actual logic part of obtaining the result, which is to parse these expressions, multiply <code>X</code> and <code>Y</code> and add the results for each of the expressions.</p>
<pre><code class="language-go">func Multiply(matches [][]string) int {
	score := 0
	for _, match := range matches {
		x, err := strconv.Atoi(string(match[1]))
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(string(match[2]))
		if err != nil {
			panic(err)
		}
		score += x * y
	}
	return score
}
</code></pre>
<p>This is pretty straightforward, we iterate over each of the matches, that is one <code>mul(X,Y)</code> expression and parse the <code>X</code> and <code>Y</code> each into integers and multiply them, the result obtained is then added to the score. We do this for each <code>mul(X,Y)</code> expression that is found in the string(line) and return the final score.</p>
<h3>Input</h3>
<p>Now, the example gave us a single string, but I realized there were six lines in the input (individual input), so we needed to parse each line and add up the score.</p>
<p>Remember this as it will be critical in part 2, it took me some time and questioning my existence to realize we need to combine all the lines to get the result (not necessary in part 1 but for sure in part 2).</p>
<h2>Part 2</h2>
<p>This is where things go wrong usually. At least for me, it did.</p>
<p>I started with a very naive approach, with a forever loop and finding the index of do or don’t and stripping off those sections, and then looping until we have no do and don’ts left. That was working for the test case but failed on the actual input.</p>
<p>The approach I finally came up and was working by tweaking the same approach slightly.</p>
<h3>Approach</h3>
<p>What I came up with is to find the first match location of <code>don’()</code> and <code>do()</code> strings in the entire string, we take that and remove the parts after <code>don’t()</code> or before <code>do()</code> . That way we can trim down the string to only valid/enabled <code>mul()</code>expressions.</p>
<p><img src="https://meetgor-cdn.pages.dev/aoc-2024-d3-solution-part2.jpg" alt="AOC Day 3 Part 2"></p>
<p>So, the approach more clearly defined will be:</p>
<ul>
<li>
<p>Find the location (index) of the <code>don’t()</code> and <code>do()</code> expressions</p>
</li>
<li>
<p>We need to keep track of whether the previous string was enabled or disabled so will keep a flag to append the enabled expressions (part of the string) to the final result</p>
</li>
<li>
<p>If none of them are found, return the string(line) as it is</p>
</li>
<li>
<p>If we found either of them then:</p>
<ul>
<li>
<p>If we find don’t first (<code>don’t()</code> appears before <code>do()</code>)</p>
<ul>
<li>
<p>If the enabled flag is true → append the string before the <code>don’t()</code> expression</p>
</li>
<li>
<p>Then toggle the enabled as false and trim off the <code>don’t()</code> part<br>
(This way we have completed checking everything before the don’t expression, so we remove that part from the line string)</p>
</li>
</ul>
</li>
<li>
<p>If we find do first (<code>do()</code> appears before <code>don’t()</code>)</p>
<ul>
<li>
<p>If the enabled flag is true → append the string before the <code>do()</code> expression</p>
</li>
<li>
<p>Then toggle the enabled as true and trim off the <code>do()</code> part<br>
(This way we have completed checking everything before the do expression, so we remove the part from the line string)</p>
</li>
</ul>
</li>
</ul>
</li>
<li>
<p>We do this until there is no line string left to be checked</p>
</li>
</ul>
<h3>Code</h3>
<p>I used simple Strings.Index to get the first match index for the substring, In this case, I want the first matching index for <code>don’t()</code> and <code>do()</code> . Once we have the indices of both the matches, we can iterate over till we are not left with any do or don’ts in the string.</p>
<p>If we have either do or don’t we append to the string the part before don’t if enabled or part before do if enabled and toggle on and off the enabled flag accordingly. By the end of the loop, the result string will have only the enabled parts of the line/string.</p>
<pre><code class="language-go">func StripDoDont(line string) string {
    result := &quot;&quot;
    enabled := true
    dontOffset := len(&quot;don't()&quot;)
    doOffset := len(&quot;do()&quot;)

    for len(line) &gt; 0 {
        dontIndex := strings.Index(line, &quot;don't()&quot;)
        doIndex := strings.Index(line, &quot;do()&quot;)

        if dontIndex == -1 &amp;&amp; doIndex == -1 {
            if enabled {
                result += line
            }
            break
        }
        
        if dontIndex != -1 &amp;&amp; (doIndex == -1 || dontIndex &lt; doIndex) {
            if enabled {
                result += line[:dontIndex]
            }
            enabled = false
            line = line[dontIndex+dontOffset:]
        } else {
            if enabled {
                result += line[:doIndex]
            }
            enabled = true
            line = line[doIndex+doOffset:]
        }
    }

    return result
}
</code></pre>
<p>I take this function and pass it to the multiply function where I get the matching patterns for the <code>mul</code> expression and do the math.</p>
<p>The <a href="https://pkg.go.dev/strings#Index">strings.Index</a> method takes in a string and a substring to find within that string and returns the index of the first occurring instance of that substring. With that we can identify if the line string contains the <code>do()</code> or <code>don’t()</code> expressions, if they don’t we simply return the line and if there are instances of them, we loop and trim the string before and after the expressions depending on whether the flag is enabled or disabled.</p>
<p>Let’s take an example and walk through the logic:</p>
<pre><code class="language-plaintext">abcxmul(1,3)don't()mul(9, 7)do()mul(1,2)don't()mul(8,7)

enabled = True
result = &quot;&quot;
line = &quot;abcxmul(1,3)don't()mul(9, 7)do()mul(1,2)don't()mul(8,7)&quot;
---
After Iteration 1:
    result -&gt; abcxmul(1,3)
    line -&gt; mul(9, 7)do()mul(1,2)don't()mul(8,7)
    enabled = False
---
After Iteration 2:
    result -&gt; abcxmul(1,3)
    line -&gt; mul(1,2)don't()mul(8,7)
    enabled = True
---
After Iteration 3:
    result -&gt; abcxmul(1,3)mul(1,2)
    line -&gt; mul(8,7)
    enabled -&gt; False
---
After Iteration 4:
    No do and don't found
    result -&gt; abcxmul(1,3)mul(1,2)
    break out of loop
---

Result -&gt; abcxmul(1,3)mul(1,2)
</code></pre>
<p>We process the result with the same <code>Multiply</code> function that we used in the first part after passing it through the <code>FindMulExpression</code> function that will return all the mul expressions in the result line string.</p>
<h3>Heads up with the input</h3>
<p>The actual input of the puzzle is I think multiple lines, so we need to preserve this state of the line in all the remaining lines. OR, we could be smarter and create a single large string and process that. Both are valid and would give the same results. I just didn’t like to add an overhead of keeping track of all the states and line, so I just concatenated all the lines and processed that single string.</p>
<h2>Conclusion</h2>
<p>This was a simple problem in essence but if you are not aware of regex you could go down a rabbit hole of writing your own parser or wired string manipulation (just like I did).</p>
<p>That’s it form day 3, I will be doing more live stream solving over the weekend and may be the next week as well. Find the code for my AoC solutions here on GitHub.</p>
<p>Till then,</p>
<p>Happy Coding :)</p>
