package kata



/*
strings ends with?

Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false
*/

func Solution(str, ending string) bool {
	if len(str) < len(ending) {
			return false
		}
	
		end := str[len(str)-len(ending):]
		return end == ending
	  
}