package kata

import "fmt"

/*
If you can't sleep, just count sheep!!

Task:
Given a non-negative integer, 3 for example,
return a string with a murmur: "1 sheep...2 sheep...3 sheep...".
Input will always be valid, i.e. no negative integers.

func countSheep(num int) string {
  // Your code here!

}
*/

func CountSheep(num int) string {
	sheepCount := ""
	for i := 1; i <= num; i++ {
		sheepCount += fmt.Sprintf("%d sheep...", i)
	}
	return sheepCount
}
