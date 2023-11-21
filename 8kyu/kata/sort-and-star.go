package kata

import (
	"sort"
	"strings"
)

/*
You will be given a list of strings. You must sort it alphabetically
(case-sensitive, and based on the ASCII values of the chars) and then return the first value.

The returned value must be a string, and have "***" between each of its letters.

You should not remove or add elements from/to the array.
*/

func TwoSort(arr []string) string {
	sort.Strings(arr) // Sort the array alphabetically
	firstWord := arr[0]

	// Join the letters of the first word with "***" in between
	return strings.Join(strings.Split(firstWord, ""), "***")
}
