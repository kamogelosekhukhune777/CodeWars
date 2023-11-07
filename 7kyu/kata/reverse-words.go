package kata 

/*
Reverse words

Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.

Examples
"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"
*/
func ReverseWords(str string) string {
    var reversedWords []rune
    var currentWord []rune

    for _, char := range str {
        if unicode.IsSpace(char) {
            reversedWords = append(reversedWords, currentWord...)
            reversedWords = append(reversedWords, ' ')
            currentWord = []rune{}
        } else {
            currentWord = append([]rune{char}, currentWord...)
        }
    }

    reversedWords = append(reversedWords, currentWord...)

    return string(reversedWords)
}





