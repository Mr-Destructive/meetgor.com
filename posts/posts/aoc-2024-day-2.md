{
  "title": "Advent of Code Day 2 in Golang: Slicing and Dicing Reports",
  "status": "published",
  "slug": "aoc-2024-day-2",
  "date": "2025-04-05T12:33:19Z"
}

<h2>Introduction</h2>
<p>So, this is day 2 of the Advent of Code 2024 in Golang, and we will be exploring my approach and solution for the same. The problem was not as easy but was pretty simple after implemented and found correct.</p>
<p>You can check out my solutions <a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day02/main.go">here on GitHub</a>.</p>
<h2>Part 1</h2>
<p>We have been given some lines called reports, and each report has a bunch of levels. and the requirement of the report is that it needs to be either strictly increasing or decreasing by a factor of at least 1 or at most 3.</p>
<p>This means if the first two elements are increasing even by one, the other subsequent elements in that report should be increasing (by 1, 2, or 3) levels, and there cannot be any change (i.e. 0 change in two adjacent numbers, or two adjacent numbers cannot be same)</p>
<pre><code class="language-plaintext">7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
</code></pre>
<ul>
<li>
<p>We first do the input parsing, it is pretty straightforward, we need to split it by <code> </code> to get individual reports, this will be a string so <code>”7 6 4 2 1&quot;</code> , we want to get a slice of integers.</p>
</li>
<li>
<p>So we go ahead and split by spaces/whitespace <code>” “</code> to get the individual levels (numbers) and we need to convert them into integers.</p>
</li>
<li>
<p>Once we have individual strings of the report i.e. levels as <code>[“7”, “6”, “4”, “2”, “1”]</code> , we need to cast them to integers.</p>
</li>
<li>
<p>We iterate over each of them and cast them to integers and append to the list.</p>
</li>
<li>
<p>Once we have constructed the list, we append to the reports list which will be the array of arrays, i.e. each line is a report, and each report has many levels so slice of slice of integers.</p>
</li>
</ul>
<pre><code class="language-go">func SplitLevels(lines []string) [][]int {
	reportLevels := [][]int{}
	for i, reportLine := range lines {
		reportLevels = append(reportLevels, []int{})
		for _, levelStr := range strings.Split(reportLine, &quot; &quot;) {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				log.Fatal(err)
			}
			reportLevels[i] = append(reportLevels[i], level)
		}
	}
	return reportLevels
}
</code></pre>
<p>Once we have constructed the reports and levels, we move ahead in actually analyzing the patterns within the levels in the individual reports.</p>
<p>For that:</p>
<ul>
<li>
<p>We first take individual reports, calculate the difference between the first two elements, and remember to use absolute difference carefully here.</p>
</li>
<li>
<p>We need to maintain a flag which indicates whether the levels in the report are increasing or decreasing, which can be determined with the first two elements.</p>
<p>That is if the first two elements are increasing, the subsequent levels should also be increasing and if they are decreasing then all the levels should be decreasing as well</p>
</li>
<li>
<p>We first have a guard check, if the difference between them is 0 or greater than 3 or less than -3 which is the condition of the levels to be safe. If that is the case then we return false that is the report is not safe.</p>
</li>
<li>
<p>We now iterate on the report after the first two elements, we then compute the difference between the next two levels, if the flag is increasing is true and the current difference is less than or equal to 0 or it exceeding 3 we also mark it as false</p>
</li>
<li>
<p>The other condition is that if the flag is is decreasing, which means the first two elements were having a negative difference, so we check if the current difference is greater than or equal to 0 or it is less than -3, if that is the case we mark that as false</p>
</li>
<li>
<p>After computing the difference for all the levels, if we come out of the loop, we return true as we didn’t see any discrepancy in the levels.</p>
</li>
</ul>
<pre><code class="language-go">

func IsSafe(report []int) (bool) {
	prevDiff := report[1] - report[0]
	isIcreasing := prevDiff &gt; 0
	if prevDiff == 0 || prevDiff &gt; 3 || prevDiff &lt; -3 {
		return false
	}

	for i := 2; i &lt; len(report); i++ {
		currDiff := report[i] - report[i-1]
		if isIcreasing {
			if currDiff &lt;= 0 || currDiff &gt; 3 {
				return false
			}
		} else {
			if currDiff &gt;= 0 || currDiff &lt; -3 {
				return false
			}
		}
	}
	return true
}
</code></pre>
<h2>Part 2</h2>
<p>For part two, we need to do a few things, we need to compute if the report is safe or not, and if that is unsafe, we can almost remove one element from the report to make it safe.</p>
<p>For that the approach is:</p>
<ul>
<li>
<p>Get the index where we first saw the discrepancy in the levels</p>
</li>
<li>
<p>Check by removing that element from the report, if that makes the report safe, then return true i.e. we found the safe report</p>
</li>
<li>
<p>If we still find the report unsafe, remove the element before the index where the discrepancy was found, if now we find it safe after removing that element, then mark it safe</p>
</li>
<li>
<p>If still we find the report unsafe, then remove the element after the index where we originally found the discrepancy, if the report becomes safe, we mark that report safe</p>
</li>
<li>
<p>Else we mark the report unsafe, as we cannot find only the element removable that makes the report safe.</p>
</li>
</ul>
<pre><code class="language-go">func RemoveAndCheck(report []int, index int) bool {
	if index &gt; len(report)-1 || index &lt; 0 {
		return false
	}
	reportNew := append([]int{}, report[:index]...)
	reportNew = append(reportNew, report[index+1:]...)
	safe, _ := IsSafe(reportNew)
	fmt.Println(safe, report)
	return safe
}

func RemoveLevels(report []int) bool {
	safe, unsafeIndex := IsSafe(report)
	if safe {
		return true
	} else {
		if RemoveAndCheck(report, unsafeIndex) {
			return true
		}
		if RemoveAndCheck(report, unsafeIndex-1) {
			return true
		}
		if RemoveAndCheck(report, unsafeIndex+1) {
			return true
		}
		return false
	}
}
</code></pre>
<p>You can check out my solutions <a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day02/main.go">here on GitHub</a>.</p>
<h2>Conclusion</h2>
<p>So that was it, a pretty simple problem for day 2 of the advent of code 2024 in Golang. Hope you enjoyed this walkthrough of the day one puzzle in the Advent of Code 2024 in Golang.</p>
<p>Let me know if you have any other interesting solutions, or if you have anything to share about this, any feedback, questions, or suggestions are welcome.</p>
<p>Thank you for reading, and I will see you tomorrow for day 3</p>
<p>Happy Coding :)</p>
