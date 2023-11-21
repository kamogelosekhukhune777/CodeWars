package kata

import (
	"strconv"
	"strings"
)

/*Acknowledgments:
I thank yvonne-liu for the idea and for the example tests :)

Description:
Encrypt this!

You want to create secret messages which can be deciphered by
the Decipher this! kata. Here are the conditions:

Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
The first letter must be converted to its ASCII code.
The second letter must be switched with the last letter
Keepin' it simple: There are no special characters in the input.
Examples:
EncryptThis("Hello") == "72olle"
EncryptThis("good") == "103doo"
EncryptThis("hello world") == "104olle 119drlo"
*/

func EncryptThis(text string) string {
	words := strings.Fields(text) // Split the text into words
	result := make([]string, len(words))

	for i, word := range words {
		if len(word) == 0 {
			continue
		}

		ascii := strconv.Itoa(int(word[0])) // Convert the first letter to ASCII code

		switch len(word) {
		case 1:
			result[i] = ascii // If the word has only one letter, append its ASCII code
		case 2:
			result[i] = ascii + string(word[1]) // If the word has two letters, swap the second letter with the ASCII code
		default:
			result[i] = ascii + string(word[len(word)-1]) + word[2:len(word)-1] + string(word[1]) // For longer words, follow the encryption rules
		}
	}

	return strings.Join(result, " ") // Join the encrypted words back into a string
}
