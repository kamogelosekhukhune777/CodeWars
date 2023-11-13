package kata

/*
Your goal in this kata is to implement a difference function, 
which subtracts one list from another and returns the result.

It should remove all values from list a, which are present in list b keeping their order.

array_diff({1, 2}, 2, {1}, 1, *z) == {2} (z == 1)
If a value is present in b, all of its occurrences must be removed from the other:

array_diff({1, 2, 2, 2, 3}, 5, {2}, 1, *z) == {1, 3} (z == 2)
*/

func ArrayDiff(a, b []int) []int {
  	result := make([]int, 0)

	// Create a map to store the elements in list b
	bMap := make(map[int]bool)
	for _, value := range b {
		bMap[value] = true
	}

	// Iterate through list a, adding elements to the result if not in list b
	for _, element := range a {
		if !bMap[element] {
			result = append(result, element)
		}
	}

	return result
}