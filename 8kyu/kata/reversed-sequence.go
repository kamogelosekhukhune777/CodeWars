package kata

/*
Build a function that returns an array of integers from n to 1 where n>0.

Example : n=5 --> [5,4,3,2,1]
*/

func ReverseSeq(n int) []int {
	result := make([]int, n)
	for i := n; i >= 1; i-- {
		result[n-i] = i
	}

	return result
}
