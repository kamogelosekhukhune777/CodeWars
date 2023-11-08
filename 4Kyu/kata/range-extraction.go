package kata

import (
	"fmt"
	"strconv"
	"strings"
)

//range Extraction
/*
A format for expressing an ordered list of integers is to use a comma separated list of either

individual integers
or a range of integers denoted by the starting integer separated from the end integer in the range
by a dash, '-'. The range includes all integers in the interval including both endpoints.
It is not considered a range unless it spans at least 3 numbers. For example "12,13,15-17"
Complete the solution so that it takes a list of integers in increasing order and returns a
correctly formatted string in the range format.

Example:

solution([-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);
// returns "-10--8,-6,-3-1,3-5,7-11,14,15,17-20"
*/

func Solution(list []int) string {
	if len(list) == 0 {
		return ""
	}

	var result []string
	rangeStart := list[0]
	rangeEnd := list[0]

	// Helper function to append a range or individual integer to the result
	appendRangeOrInteger := func(start, end int) {
		if start == end {
			result = append(result, strconv.Itoa(start))
		} else if end-start >= 2 {
			result = append(result, fmt.Sprintf("%d-%d", start, end))
		} else {
			for i := start; i <= end; i++ {
				result = append(result, strconv.Itoa(i))
			}
		}
	}

	// Iterate through the list to find and format ranges
	for i := 1; i < len(list); i++ {
		if list[i] == list[i-1]+1 {
			rangeEnd = list[i]
		} else {
			appendRangeOrInteger(rangeStart, rangeEnd)
			rangeStart = list[i]
			rangeEnd = list[i]
		}
	}

	// Append the last range or individual integer
	appendRangeOrInteger(rangeStart, rangeEnd)

	return strings.Join(result, ",")
}
