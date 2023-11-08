package kata



/*MOVING ZEROS TO THE END

Write an algorithm that takes an array and moves all of the zeros to the end, preserving the order of the other elements.

MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }
*/

func MoveZeros(arr []int) []int {
	// Create a variable to keep track of the current index where non-zero elements should be placed
	nonZeroIndex := 0

	// Iterate through the array
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			// If the element is non-zero, move it to the current nonZeroIndex
			arr[nonZeroIndex], arr[i] = arr[i], arr[nonZeroIndex]
			nonZeroIndex++
		}
	}

	return arr
}