package kata

/*
Lyrics...
Pyramids are amazing! Both in architectural and mathematical sense.
If you have a computer, you can mess with pyramids even if you are
not in Egypt at the time. For example, let's consider the following
problem. Imagine that you have a pyramid built of numbers, like this one here:

   /3/
  \7\ 4
 2 \4\ 6
8 5 \9\ 3
Here comes the task...
Let's say that the 'slide down' is the maximum sum of consecutive numbers
from the top to the bottom of the pyramid. As you can see, the longest
'slide down' is 3 + 7 + 4 + 9 = 23

Your task is to write a function that takes a pyramid representation as an argument
and returns its largest 'slide down'. For example:

* With the input `[[3], [7, 4], [2, 4, 6], [8, 5, 9, 3]]`
* Your function should return `23`.
By the way...
My tests include some extraordinarily high pyramids so as you can guess, brute-force
method is a bad idea unless you have a few centuries to waste. You must come up with
something more clever than that.

(c) This task is a lyrical version of Problem 18 and/or Problem 67 on ProjectEuler.
*/

func LongestSlideDown(pyramid [][]int) int {
	// Create a 2D slice to store the maximum path sums
	dp := make([][]int, len(pyramid))
	for i := range dp {
		dp[i] = make([]int, len(pyramid[i]))
	}

	// Initialize the bottom row of the dp array with the values from the bottom row of the pyramid
	copy(dp[len(pyramid)-1], pyramid[len(pyramid)-1])

	// Traverse the pyramid from bottom to top and calculate the maximum path sums
	for row := len(pyramid) - 2; row >= 0; row-- {
		for col := 0; col < len(pyramid[row]); col++ {
			// Calculate the maximum path sum for the current position
			dp[row][col] = pyramid[row][col] + max(dp[row+1][col], dp[row+1][col+1])
		}
	}

	// The top of the dp array will contain the maximum path sum
	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
