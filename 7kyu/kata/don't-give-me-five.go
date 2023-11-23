package kata

//Don't give me five!
/*
In this kata you get the start number and the end number of a region and should return the
count of all numbers except numbers with a 5 in it. The start and the end number are both inclusive!

Examples:

1,9 -> 1,2,3,4,6,7,8,9 -> Result 8
4,17 -> 4,6,7,8,9,10,11,12,13,14,16,17 -> Result 12
The result may contain fives. ;-)
The start number will always be smaller than the end number. Both numbers can be also negative!
*/

func DontGiveMeFive(start int, end int) int {
	count := 0

	for num := start; num <= end; num++ {
		if !containsFive(num) {
			count++
		}
	}

	return count
}

func containsFive(num int) bool {
	if num < 0 {
		num *= -1 // Convert negative numbers to positive for digit checking
	}

	for num > 0 {
		if num%10 == 5 {
			return true
		}
		num /= 10
	}

	return false
}
