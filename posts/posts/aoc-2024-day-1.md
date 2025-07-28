{
  "title": "Advent of Code, 2024, Day 1 in Golang: Historian Hysteria",
  "status": "published",
  "slug": "aoc-2024-day-1",
  "date": "2025-04-05T12:33:19Z"
}

<h2>Introduction</h2>
<p>Hello everyone, it’s that time of the year, Advent of Code, I will be solving this year as well with Golang. In previous years I have been doing Advent of Code but was not able to keep up with the pace and left it midway (not even halfway). This year however I am determined and want to solve all the problems as much as I can.</p>
<p>Let’s dive into the first day which should be and is pretty simple and straightforward. A simple list and map creation and traversal and basic math operations.</p>
<p>I also live-streamed the solution, you can check it out the <a href="https://www.youtube.com/live/3K02tEEBgto?si=ojS5rsh5nGpk3U-B">stream on YouTube</a></p>
<p>And also a <a href="https://youtu.be/4U97gLyz0Ss?si=SvINHaGz-mow_q3O">shorter video</a> on the approach and solution in Golang.</p>
<!-- raw HTML omitted -->
<p>Or you stick here and continue reading. Thank you</p>
<p>You can check out my solutions <a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day01/main.go">here on GitHub</a>.</p>
<h2>Part 1</h2>
<p><a href="https://adventofcode.com/2024/day/1">Advent of code, 2024, day 1</a></p>
<p>We are given two lists here, the first part aims to find the absolute difference (distance) between each element sorted from smallest to largest.</p>
<p>So, in essence, we take the two lists, sort them and one by one, for each corresponding element paired up, we take the absolute difference and sum that difference up for all the numbers in the list.</p>
<pre><code class="language-plaintext">3   4
4   3
2   5
1   3
3   9
3   3
</code></pre>
<p>So, first, we need to split the input into different lists:</p>
<ol>
<li>
<p>We first range over all the lines, initialize two empty lists of integers</p>
</li>
<li>
<p>Then we split the line with the space as the separator, so this gives us the slice of strings as <code>[“3”, “4”]</code></p>
</li>
<li>
<p>But we need to elements as integers, so take the first number and convert it to integer, similarly we do it for the second number.</p>
</li>
<li>
<p>Then once we have both numbers, we append them to the corresponding lists, the first number goes to the first list, and the second is appended to the second list.</p>
</li>
<li>
<p>Then we return the two lists</p>
</li>
</ol>
<p>NOTE: You cannot take the difference of those two numbers here itself, since we need to find the smallest number and sort the numbers in each list, so we need to get the lists populated first.</p>
<pre><code class="language-go">func SplitLists(lines []string) ([]int, []int) {
	listOne := []int{}
	listTwo := []int{}

	for _, line := range lines {
		// |3   4
		// [&quot;3&quot;,&quot;4&quot;] slice of string ([]string)
		// 3 
        // 4
        // [3,4] slice of int ([]int)
		numbers := strings.Split(line, &quot;   &quot;)
		numOne, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		numTwo, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		listOne = append(listOne, numOne)
		listTwo = append(listTwo, numTwo)
	}
	return listOne, listTwo
}
</code></pre>
<p>In the above code, we have initialized two empty slices of strings, we take the parameter lines which is a slice of string, indicating a line-by-line string representation of the input. I have parsed the input with these helper functions.</p>
<p>The ReadFileBytes and ReadFileLines, one the bytes, the other gives the line-by-line string which gives a slice of strings.</p>
<p>So once we have the lines, we iterate over each line and split the lines on space to get the two numbers. So, the line “<code>3 4</code>“ will be split into <code>[“3”, “4”]</code> . Now we get the first element and convert it into an integer as we need to sort and take the difference later.</p>
<p>So, by accessing the first and second elements in the split line as <code>numbers[0]</code> and <code>numbers[1]</code> and converting the type to integer, <a href="https://pkg.go.dev/strconv#Atoi">strconv.Atoi</a> function, which takes in a string and gives back an integer or an error.</p>
<p>Now, we have two numbers as integers, we append the first element to the first element as <code>listOne = append(listOne, numOne)</code> and <code>listTwo = append(listTwo, numTwo)</code></p>
<p>So, we append one by one as we iterate over the input through all the lines, so at the end of this function, we will have two lists of integers.</p>
<pre><code>[3 4 2 1 3 3]
[4 3 5 3 9 3]
</code></pre>
<p>Then, once we have the slices of integers, we sort those lists. Then we range over the lists one by one element, since both the lists are of the same size, we can reference one by the index of the other.</p>
<p>Then for each difference of the two integers (one from the first list and the other from the second list), we cast it to a <code>float64</code> and pass it to the <a href="https://pkg.go.dev/math#Abs">math.Abs</a> function, which is annoying as Golang doesn’t have an absolute function for intgers. We cast the integer to float for parsing it to the Abs method and cast the returned float64 value back to int. Kind of wired but fine.</p>
<p>We keep adding the absolute differences for each paired difference of the elements in the two lists. At the end, we will have a final score which is the score for part one.</p>
<pre><code class="language-go">func PartOne(lines []string) int {
	listOne, listTwo := SplitLists(lines)
	sort.Ints(listOne)
	sort.Ints(listTwo)
	totalScore := 0
	for i := range listOne {
		totalScore += int(math.Abs(float64(listOne[i] - listTwo[i]))
	}
	return totalScore
}
</code></pre>
<h2>Part 2</h2>
<p>For part two, we need to take all the numbers in the first list count the number of times that number has occurred in the second list, and take a product of them and add it up for all the numbers.</p>
<p>So in the example:</p>
<pre><code class="language-plaintext">3   4
4   3
2   5
1   3
3   9
3   3
</code></pre>
<p>The numbers in the first list are [3,4,2,1,3,3]</p>
<p>We have to count the occurrences of each of them in the second list</p>
<p>So, in the second list [4,3,5,3,9,3], the number <code>3</code> occurs <code>3</code> times, so we do <code>3×3</code> which is <code>9</code> and then, do the same for <code>4</code> which occurs only once in the second list so, we get <code>4</code>, then <code>2</code> occurs <code>0</code> times, so we get <code>0</code></p>
<p>We get → <code>(3×3) + (4×1) + (2×0) + (1×0) + (3×3) + (3×3)</code></p>
<p>The first number is the element in the first list and the second number is the occurrence of that number in the second list.</p>
<p>which comes out to be <code>9+4+0+0+9+9</code> , so the answer is <code>31</code> for the example.</p>
<p>Once it is clear, what we have to do, we simply have to iterate over the second list and create a map of the frequency/occurances/number of times it appears in that list.</p>
<h3>Solution</h3>
<p>So, we will have to modify the <code>SplitLists</code> functions a bit, we need to split and also map the second list with the key as the number itself and the value as its count in the second list.</p>
<p>Just that change, we create an additional return value with an empty map of integers with integers. The mapTwo variable will be a map that will have a key as the number in the second list and its value as the number of times it is present in that list.</p>
<pre><code class="language-go">func SplitListsAndMap(lines []string) ([]int, map[int]int) {
	listOne := []int{}
	listTwoCounts := make(map[int]int)

	for _, line := range lines {
		numbers := strings.Split(line, &quot;   &quot;)
		numOne, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		numTwo, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		listOne = append(listOne, numOne)
        listTwoCounts[numTwo] += 1
	}
	return listOne, listTwoCounts
}
</code></pre>
<p>So, as we iterate over each line, we parse the string number into an integer and increment its count in the map.</p>
<pre><code>[3 4 2 1 3 3]
map[3:3 4:1 5:1 9:1]
</code></pre>
<p>In the actual calculation of the score, we need to iterate over the elements of the first list and multiply the number with its count in the second list as we now have the map of it. We multiply those and add them up for each line, which becomes the final score.</p>
<pre><code class="language-go">func PartTwo(lines []string) int {
    similarityScore := 0

	listOne, mapTwo := SplitListsAndMap(lines)

	for _, numOne := range listOne {
		score := numOne * mapTwo[numOne]
		similarityScore += score
	}

	return similarityScore
}
</code></pre>
<p>So, that is how we got the final score for part two.</p>
<p>You can check out my solutions <a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day01/main.go">here on GitHub</a>.</p>
<h2>Conclusion</h2>
<p>So that was it, a pretty simple problem for day 1 of the advent of code 2024 in Golang. Hope you enjoyed this walkthrough of the day one puzzle in the Advent of Code 2024 in Golang.</p>
<p>Thank you for reading, and I will see you tomorrow for day 2</p>
<p>Happy Coding :)</p>
