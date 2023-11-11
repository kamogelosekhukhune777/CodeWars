package kata

import "strconv"

//reverse or rotate
/*
The input is a string str of digits. Cut the string into chunks
(a chunk here is a substring of the initial string) of size sz
(ignore the last chunk if its size is less than sz).

If a chunk represents an integer such as the sum of the cubes of its
digits is divisible by 2, reverse that chunk; otherwise rotate it to the left
by one position. Put together these modified chunks and return the result as a string.

If

sz is <= 0 or if str is empty return ""
sz is greater (>) than the length of str it is impossible to take a chunk of size sz hence return "".
Examples:
revrot("123456987654", 6) --> "234561876549"
revrot("123456987653", 6) --> "234561356789"
revrot("66443875", 4) --> "44668753"
revrot("66443875", 8) --> "64438756"
revrot("664438769", 8) --> "67834466"
revrot("123456779", 8) --> "23456771"
revrot("", 8) --> ""
revrot("123456779", 0) --> ""
revrot("563000655734469485", 4) --> "0365065073456944"
Example of a string rotated to the left by one position:
s = "123456" gives "234561".
*/

func Revrot(str string, sz int) string {
	if sz <= 0 || sz > len(str) || str == "" {
		return ""
	}

	chunks := make([]string, len(str)/sz)
	for i := 0; i < len(chunks); i++ {
		chunk := str[i*sz : (i+1)*sz]
		if isSumOfCubesEven(chunk) {
			chunks[i] = reverse(chunk)
		} else {
			chunks[i] = rotateLeft(chunk)
		}
	}

	return concat(chunks)
}

func isSumOfCubesEven(chunk string) bool {
	sum := 0
	for _, digit := range chunk {
		num, _ := strconv.Atoi(string(digit))
		sum += num * num * num
	}
	return sum%2 == 0
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func rotateLeft(s string) string {
	return s[1:] + string(s[0])
}

func concat(chunks []string) string {
	result := ""
	for _, chunk := range chunks {
		result += chunk
	}
	return result
}
