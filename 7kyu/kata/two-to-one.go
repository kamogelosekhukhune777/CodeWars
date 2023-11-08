package kata

import "sort"

/*
Two to One

Take 2 strings s1 and s2 including only letters from a to z.
Return a new sorted string, the longest possible,
containing distinct letters - each taken only once - coming from s1 or s2.

Examples:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"
*/

func TwoToOne(s1 string, s2 string) string {
	// Combine the two strings
	combined := s1 + s2

	// Remove duplicate characters by creating a map of unique characters
	uniqueChars := make(map[rune]bool)
	for _, char := range combined {
		uniqueChars[char] = true
	}

	// Create a slice to store the unique characters
	var uniqueCharsList []rune
	for char := range uniqueChars {
		uniqueCharsList = append(uniqueCharsList, char)
	}

	// Sort the unique characters
	sort.Slice(uniqueCharsList, func(i, j int) bool {
		return uniqueCharsList[i] < uniqueCharsList[j]
	})

	// Convert the sorted unique characters back to a string
	result := string(uniqueCharsList)

	return result
}
