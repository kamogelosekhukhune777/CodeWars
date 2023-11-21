package kata

import "strings"

/*
Create two functions to encode and then decode a string using the Rail Fence Cipher. This cipher
is used to encode a string by placing each character successively in a diagonal along a set of "rails".
First start off moving diagonally and down. When you reach the bottom, reverse direction and move diagonally
and up until you reach the top rail. Continue until you reach the end of the string. Each "rail" is then r
ead left to right to derive the encoded string.

For example, the string "WEAREDISCOVEREDFLEEATONCE" could be represented in a three rail system as follows:

W       E       C       R       L       T       E
  E   R   D   S   O   E   E   F   E   A   O   C
    A       I       V       D       E       N
The encoded string would be:

WECRLTEERDSOEEFEAOCAIVDEN
Write a function/method that takes 2 arguments, a string and the number of rails,
and returns the ENCODED string.

Write a second function/method that takes 2 arguments, an encoded string and the number
of rails, and returns the DECODED string.

For both encoding and decoding, assume number of rails >= 2 and that passing an empty string
 will return an empty string.

Note that the example above excludes the punctuation and spaces just for simplicity. There are,
 however, tests that include punctuation. Don't filter out punctuation as they are a part of the string.
*/

func Encode(s string, n int) string {
	if n <= 1 || len(s) <= 1 {
		return s
	}

	rails := make([]string, n)
	row := 0
	down := true

	for _, char := range s {
		rails[row] += string(char)

		if down {
			row++
			if row == n {
				row -= 2
				down = false
			}
		} else {
			row--
			if row == -1 {
				row += 2
				down = true
			}
		}
	}

	return strings.Join(rails, "")
}

func Decode(s string, n int) string {
	if n <= 1 || len(s) <= 1 {
		return s
	}

	rails := make([][]rune, n)
	for i := range rails {
		rails[i] = make([]rune, len(s))
	}

	row := 0
	down := true
	index := 0

	for i := 0; i < len(s); i++ {
		rails[row][i] = rune(s[index])
		index++

		if down {
			row++
			if row == n {
				row -= 2
				down = false
			}
		} else {
			row--
			if row == -1 {
				row += 2
				down = true
			}
		}
	}

	decoded := ""
	row = 0
	down = true

	for i := 0; i < len(s); i++ {
		decoded += string(rails[row][i])

		if down {
			row++
			if row == n {
				row -= 2
				down = false
			}
		} else {
			row--
			if row == -1 {
				row += 2
				down = true
			}
		}
	}

	return decoded
}
