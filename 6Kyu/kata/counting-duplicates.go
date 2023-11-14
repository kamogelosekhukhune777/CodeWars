package kata

import (
	"strings"
	"unicode"
)

/*
Counting Duplicates::

Count the number of Duplicates:
Write a function that will return the count of distinct case-insensitive
alphabetic characters and numeric digits that occur more than once in the
input string. The input string can be assumed to contain only alphabets
(both uppercase and lowercase) and numeric digits.

Example
"abcde" -> 0 # no characters repeats more than once
"aabbcde" -> 2 # 'a' and 'b'
"aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
"indivisibility" -> 1 # 'i' occurs six times
"Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
"aA11" -> 2 # 'a' and '1'
"ABBA" -> 2 # 'A' and 'B' each occur twice
*/

func Duplicate_count(s1 string) int {
	// Create a map to store the frequency of characters
	charFrequency := make(map[rune]int)

	// Convert the input string to lowercase to make the comparison case-insensitive
	s1 = strings.ToLower(s1)

	// Iterate over the characters in the input string
	for _, char := range s1 {
		// Check if the character is an alphabet letter or a numeric digit
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			charFrequency[char]++
		}
	}

	// Count the characters that occur more than once
	count := 0
	for _, freq := range charFrequency {
		if freq > 1 {
			count++
		}
	}

	return count
}
