package kata

import (
	"strings"
	"unicode"
)

/*
Given a string, capitalize the letters that occupy even indexes
and odd indexes separately, and return as shown below. Index 0
will be considered even.

For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF'].
See test cases for more examples.

The input will be a lowercase string with no spaces.
*/

func Capitalize(st string) []string {
	evenResult := strings.Builder{}
	oddResult := strings.Builder{}

	for i, char := range st {
		if i%2 == 0 {
			evenResult.WriteRune(unicode.ToUpper(char))
			oddResult.WriteRune(char)
		} else {
			evenResult.WriteRune(char)
			oddResult.WriteRune(unicode.ToUpper(char))
		}
	}

	return []string{evenResult.String(), oddResult.String()}
}
