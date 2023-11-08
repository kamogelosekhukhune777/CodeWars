package kata

/*
Implement a function that accepts 3 integer values
a, b, c. The function should return true if a triangle
can be built with the sides of given length and false in
any other case.

(In this case, all triangles must have surface greater
than 0 to be accepted).
*/
func IsThisDigits(a, b, c int) bool {
	// Check the triangle inequality theorem
	if a+b > c && a+c > b && b+c > a {
		// All sides satisfy the inequality, and the triangle is possible
		return true
	}
	// The sides do not form a valid triangle
	return false
}
