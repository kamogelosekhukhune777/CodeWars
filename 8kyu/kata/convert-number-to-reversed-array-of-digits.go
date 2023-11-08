package kata

import "strconv"

/*Convert number to reversed array of digits
Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

Example(Input => Output):
35231 => [1,3,2,5,3]
0 => [0]*/

func Digitize(n int) []int {
	// Convert the number to a string
	str := strconv.Itoa(n)

	// Reverse the string
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	// Convert each character in the reversed string to an integer and store in a slice
	reversedDigits := make([]int, len(reversedStr))
	for i, char := range reversedStr {
		digit, _ := strconv.Atoi(string(char))
		reversedDigits[i] = digit
	}

	return reversedDigits
}
