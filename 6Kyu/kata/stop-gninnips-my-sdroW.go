package kata

import "strings"

/*
Write a function that takes in a string of one or more words,
and returns the same string, but with all five or more letter
words reversed (Just like the name of this Kata). Strings passed
in will consist of only letters and spaces. Spaces will be included
only when more than one word is present.

Examples:

spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
spinWords( "This is a test") => returns "This is a test"
spinWords( "This is another test" )=> returns "This is rehtona test"
func SpinWords(str string) string {
  return "pizza"
}// SpinWords
*/

func SpinWords(str string) string {
	words := strings.Fields(str) // Split the string into words
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverses(word) //reverse
		}
	}
	return strings.Join(words, " ")
}

func reverses(s string) string { //reverses => reverse
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
