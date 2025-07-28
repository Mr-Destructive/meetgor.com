{
  "title": "Advent of Code Day 4 in Golang: Finding XMAS and X-MAS",
  "status": "published",
  "slug": "aoc-2024-day-4",
  "date": "2025-04-05T12:33:18Z"
}

<h2>Introduction</h2>
<p>Moving on to day 4, we have a grid problem in front of us, we are given some numbers in the form of a grid, i.e. some rows and columns with some upper case letters. What we need to do is to find is the word <code>XMAS</code> in any direction (up, left, down, right, diagonals), and in the second part we need to find the word <code>MAS</code> forming an X.</p>
<p>So, let’s see how we can approach this and solve it in golang.</p>
<p>You can check out my solutions <a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day04/main.go">here on GitHub</a><a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day04/main.go">.</a></p>
<h3><a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day02/main.go">Constructin</a>g the grid</h3>
<p>The most fundamental part of the problem lies in actually converting the text into a grid or a matrix form. We can split the lines, into individual lines and append each character as an element in a list, and that way we can have a list of list of strings which is a matrix or grid-like (2-dimensional) structure.</p>
<p>So, below is the input for the puzzle.</p>
<pre><code class="language-plaintext">MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
</code></pre>
<p>We need to convert it into something like this</p>
<pre><code class="language-plaintext">[
    [M M M S X X M A S M]
    [M S A M X M S M S A]
    [A M X S X M A A M M]
    [M S A M A S M S M X]
    [X M A S A M X A M M]
    [X X A M M X X A M A]
    [S M S M S A S X S S]
    [S A X A M A S A A A]
    [M A M M M X M M M M]
    [M X M X A X M A S X]
]
</code></pre>
<p>So, this is a list of strings, we can say in golang it is a <code>[][]string</code> . We can do that by creating a function like this:</p>
<pre><code class="language-go">func ConstructGrid(lines []string) [][]string {
	grid := [][]string{}
	for _, line := range lines {
		row := []string{}
		for _, char := range strings.Split(line, &quot;&quot;) {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	return grid
}
</code></pre>
<p>The above function takes in a list of strings and returns a list of list of strings that are individual letters in the grid.<br>
We can read the file bytes and split the bytes on newline characters and then this will be used as the input for this function.</p>
<p>So, once the input is parsed into a grid, we can start thinking about the actual logic of finding the word <code>XMAS</code> in it.</p>
<h2>Part 1</h2>
<p>So, in the first part, we need to find the word <code>XMAS</code> in the matrix which could be appearing:</p>
<ul>
<li>
<p>forwards (as <code>XMAS</code>)</p>
</li>
<li>
<p>backward (as <code>SAMX</code>)</p>
</li>
<li>
<p>upwards</p>
<pre><code class="language-plaintext">       S
       A
       M
       X
   ```

</code></pre>
</li>
<li>
<p>downwards</p>
<pre><code class="language-plaintext">       X
       M
       A
       S
   ```

</code></pre>
</li>
<li>
<p>Diagonal upwards (right or up left)</p>
<pre><code class="language-plaintext">       S
         A
           M
             X

       OR
             S
           A
         M 
       X
   ```

</code></pre>
</li>
<li>
<p>Diagonals downwards (right or left)</p>
<pre><code class="language-plaintext">                X
              M
            A
          S

       OR
       X
         M
           A
             S
   ```


</code></pre>
</li>
</ul>
<p>So, there are 8 directions where <code>XMAS</code> could appear in the grid, there could n number of these <code>XMAS</code> . We need to find the count of these in the grid.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1733761379973/8a0a0263-b286-47f1-a338-410dec2c6b7b.png" alt=""></p>
<p>To approach this, we can either find the first character in the word <code>XMAS</code> and then search in all directions one by one and check if we find <code>M</code> and if we have found <code>M</code> in any of the direction, we keep moving ahead in that direction and check if there is <code>A</code> and <code>S</code> in that direction.</p>
<p>The approach looks like this:</p>
<ul>
<li>
<p>Initialize the counter to 0</p>
</li>
<li>
<p>Iterate over each line</p>
<ul>
<li>
<p>Iterate over each character in the line</p>
<ul>
<li>
<p>Check if the character is equal to <code>X</code></p>
</li>
<li>
<p>If the character is <code>X</code>→</p>
<ul>
<li>
<p>Iterate over all the directions (up, down, right, left, up-left, up-right, down-left, down-right)</p>
<ul>
<li>
<p>For that direction if we find the character to be <code>M</code></p>
</li>
<li>
<p>Keep moving ahead in the same direction to find <code>A</code> and <code>S</code> similarly, if we found all the characters <code>XMAS</code> then, increment the counter</p>
</li>
<li>
<p>Else choose another direction in the loop</p>
</li>
</ul>
</li>
</ul>
</li>
</ul>
</li>
</ul>
</li>
</ul>
<p>This looks complex and large but is simple, focus one thing at a time and you can solve this easily.</p>
<p>So, for the implementation of this, we need to define a few things first:</p>
<pre><code class="language-go"> var directions [][]int = [][]int{
	[]int{0, -1},   // up
	[]int{0, 1},  // down
	[]int{1, 0},   // right
	[]int{-1, 0},  // left
	[]int{1, -1},   // up right
	[]int{-1, -1},  // up left
	[]int{1, 1},  // down right
	[]int{-1, 1}, // down left
}

var wordList []string = []string{&quot;X&quot;, &quot;M&quot;, &quot;A&quot;, &quot;S&quot;}
</code></pre>
<p>So, we have defined the list of integers in the directions which are the x and y coordinates that we need to add or subtract to get to the desired location. It is basically like a unit vector, it has a distance of 1 and a direction indicated by <code>+</code> or <code>-</code> indicating to move the left or right for x coordinates and up and down for y c-ordinates.</p>
<p>So, let me explain that more clearly, let’s say I am at <code>(1,2)</code> in the grid which is of 4x4 dimension.</p>
<pre><code class="language-plaintext">A B C D
E F G H
I J K L
M N O P
</code></pre>
<p>So, at 2,1 we have <code>G</code> , so we check some directions for this</p>
<p>up → <code>0,-1</code> → 2+0, 1-1 → 2,0, we have moved to <code>C</code></p>
<p>right → <code>1,0</code> → 2+1, 1+0 → 3,1 , we have moved to <code>H</code></p>
<p>down, left → <code>-1,1</code> → 2-1, 1+1 → 1, 2, we have moved to <code>J</code></p>
<p>So, you get the idea, that we are moving in some directions using these coordinates.</p>
<p>We can use these to get the next direction jump we want to make to find if that element has the next character in the word that we are searching.</p>
<p>We will write a function that does this first and abstract the function that checks if we have found the word in the grid.</p>
<pre><code class="language-go">func TraverseGrid(grid [][]string) int {
	score := 0
	for x, row := range grid {
		for y, char := range row {
			if char == wordList[0] {
				for _, direction := range directions {
					if FindXMAS(x, y, 1, direction, grid) {
						score += 1
					}
				}
			}
		}
	}
	return score
}
</code></pre>
<p>The above function takes in a grid and returns an integer which will be the score i.e. the count of words <code>XMAS</code> found in the grid/matrix.</p>
<p>First, we need to iterate through each of the rows in the grid, for each row, we iterate over the character, so we will have x and y coordinates as the index of the grid. We need to then check if the current character is <code>X</code> or <code>wordList[0]</code> , if that is the case, we iterate over all the directions and check if we can find <code>XMAS</code> i.e. <code>MAS</code> in that direction, if so we increment the counter. What is the <code>FindXMAS</code> function, let’s abstract that away, and pass in the <code>x</code>, <code>y</code>, which are the coordinates of the current word, <code>1</code> which will be the word position of the <code>XMAS</code> and in this case, we already have found <code>X</code> we need to find <code>MAS</code> in that direction. We pass the grid and the direction, so this function will return true or false if that direction has <code>MAS</code> in it.</p>
<p>So to iterate:</p>
<ul>
<li>
<p>We iterate over the grid and get <code>row</code> and <code>x</code> as the list of strings and index of the current row.</p>
</li>
<li>
<p>For each row i.e. list of strings, we iterate over the list of strings to get <code>char</code> and <code>y</code> as the character (string) and the index of that character in the list of the string.</p>
</li>
<li>
<p>If we find the current character to be equal to <code>X</code> which is the 0th index of the <code>wordList</code> then</p>
<ul>
<li>
<p>We iterate over all the directions and call the function <code>FindXMAS</code> to check if the remaining word <code>MAS</code> in that direction</p>
</li>
<li>
<p>If we find all the words, we increment the counter.</p>
</li>
</ul>
</li>
<li>
<p>So, we return the counter as we count the number of words <code>XMAS</code> in the grid/matrix.</p>
</li>
</ul>
<p>Now, we can implement the <code>FindXMAS</code> function, that takes in a <code>x</code>, <code>y</code> coordinates, the <code>wordPosition</code>, the direction and the grid, and return if the word is found.</p>
<ul>
<li>
<p>First, we take the current x coordinates and add the direction’s x component (0th index or first element)</p>
</li>
<li>
<p>add current y coordinates to the direction’s y component (1st index or second element)</p>
</li>
<li>
<p>if the word position i.e.. the word index or the word itself in the current function is equal to the wordList, it means that it has found the required word completely</p>
</li>
<li>
<p>We need to check by adding the direction to the x and y coordinates, we are not overshooting the width and height of the grid, so if we do, we return a false</p>
</li>
<li>
<p>The final if is for checking if the current character is equal to the word that we are looking for, it could be <code>M</code>, <code>A</code> , or <code>S</code> . If so, we return the recursively call the <code>FindXMAS</code> function by passing the updated x and y coordinates and the next word in the wordList, we keep the direction the same and pass the entire grid.</p>
</li>
</ul>
<pre><code class="language-go">func FindXMAS(x, y, wordPosition int, direction []int, grid [][]string) bool {
	xNext := x + direction[0]
	yNext := y + direction[1]
	if wordPosition &gt; len(wordList)-1 {
		return true
	}

	if xNext &lt; 0 || xNext &gt;= len(grid) || yNext &lt; 0 || yNext &gt;= len(grid[x]) {
		return false
	}

	if grid[xNext][yNext] == wordList[wordPosition] {
		return FindXMAS(xNext, yNext, wordPosition+1, direction, grid)
	}
	return false

}
</code></pre>
<p>So, we have implemented the <code>FindXMAS</code> function, this will just return if we have found the <code>MAS</code> word by going in a particular direction by updating the coordinates and checking if the word at that position in the gird is the next word in <code>MAS</code> list.</p>
<p>So, this is what the entire first part looks like:</p>
<pre><code class="language-go">func main() {
	lines := ReadFileLines(&quot;sample.txt&quot;)
	grid := ConstructGrid(lines)
	score := TraverseGrid(grid)
	fmt.Println(score)
}
</code></pre>
<p>We take in the lines as a list of strings and pass that to <code>ConstructGrid</code> and get the grid, finally, we call <code>TraverseGrid</code> , by passing the grid and getting the score as the count of the words <code>XMAS</code> in the grid.</p>
<p>That’s it from the part 1.</p>
<h2>Part 2</h2>
<p>For part two, we need to find <code>MAS</code> in the cross shape, like below:</p>
<pre><code class="language-plaintext">M.S
.A.
M.S
</code></pre>
<p>So, to solve this, we can do a similar approach but much simpler, we just need to find <code>A</code> as there will always be the word <code>MAS</code> in the center, so we just check if we have <code>A</code> and the top-left, top-right, or bottom-right, bottom-left has <code>M</code> or <code>S</code> .</p>
<p>We get the coordinates of the top-right, top-left positions, down-right, and down-left positions by adding and subtracting 1 from it. We make a basic check if we are not overshooting the boundary of the grid. If we overshoot the boundaries, we won’t find the <code>MAS</code></p>
<p>But if we are within the grid, we now get the character at those 4 positions, we check if the top-left and bottom-right have <code>M</code> and <code>S</code> or <code>S</code> or <code>M</code>, similarly for top-right and bottom-left has <code>M</code> and <code>S</code> or <code>S</code> or <code>M</code> respectively. This is the diagonal search for <code>M</code> and <code>S</code> above and below the character <code>A</code>.</p>
<p>So, if we have both the diagonal matched we return true.</p>
<pre><code class="language-go">

func FindMAS(x, y int, grid [][]string, wordList []string) bool {
	xL, yT := x-1, y+1 // Top-left neighbor
	xR, yD := x+1, y-1 // Bottom-right neighbor

	// Check if the indices are within bounds
	if xL &lt; 0 || xR &gt;= len(grid) || yT &lt; 0 || yD &lt; 0 ||
		yT &gt;= len(grid[xL]) || yD &gt;= len(grid[xR]) {
		return false
	}

	topLeft := grid[xL][yT]
	bottomRight := grid[xR][yD]
	topRight := grid[xR][yT]
	bottomLeft := grid[xL][yD]

	word1, word3 := wordList[1], wordList[3]

	isDiagonalMatch := (topLeft == word1 &amp;&amp; bottomRight == word3) || (topLeft == word3 &amp;&amp; bottomRight == word1)
	isAntiDiagonalMatch := (topRight == word1 &amp;&amp; bottomLeft == word3) || (topRight == word3 &amp;&amp; bottomLeft == word1)

	return isDiagonalMatch &amp;&amp; isAntiDiagonalMatch
}
</code></pre>
<p>So, that is the simple implementation for finding the <code>MAS</code> the diagonal.</p>
<p>Now, we need to change the <code>TraverseGrid</code> a bit, as we just iterate over the grid, and check if we have <code>A</code> in the character in the row, i.e. <code>wordList[2]</code>. Now, if we have <code>A</code>, we need to call the <code>FindMAS</code> function with the current coordinates and the grid, if that function returns true, we increment the counter,.</p>
<pre><code class="language-go">
func TraverseGrid2(grid [][]string) int {
	score := 0
	for x, row := range grid {
		for y, char := range row {
			if char == wordList[2] {
				if FindMAS(x, y, grid) {
					score += 1
				}

			}
		}
	}
	return score
}
</code></pre>
<p>So, that is the final implementation of part 2, we get the count of <code>MAS</code> in the cross direction.</p>
<p>You can check out my solutions <a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day04/main.go">here on GitHub</a><a href="https://github.com/Mr-Destructive/advent_of_code/blob/main/2024/src/day04/main.go">.</a></p>
<h2>Conclusion</h2>
<p>So, that is it from day 4 of Advent of Code in Golang, let me know if you have any suggestions, and how you approached it. any better solutions?</p>
<p>Happy Coding :)</p>
