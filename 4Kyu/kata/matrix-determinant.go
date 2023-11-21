package kata

/*
Write a function that accepts a square matrix (N x N 2D array) and
returns the determinant of the matrix.

How to take the determinant of a matrix -- it is simplest to start
with the smallest cases:

A 1x1 matrix |a| has determinant a.

A 2x2 matrix [ [a, b], [c, d] ] or

|a  b|
|c  d|
has determinant: a*d - b*c.

The determinant of an n x n sized matrix is calculated by reducing the
problem to the calculation of the determinants of n matrices ofn-1 x n-1 size.

For the 3x3 case, [ [a, b, c], [d, e, f], [g, h, i] ] or

|a b c|
|d e f|
|g h i|
the determinant is: a * det(a_minor) - b * det(b_minor) + c * det(c_minor)
where det(a_minor) refers to taking the determinant of the 2x2 matrix created
by crossing out the row and column in which the element a occurs:

|- - -|
|- e f|
|- h i|
Note the alternation of signs.

The determinant of larger matrices are calculated analogously, e.g. if M is a 4x4
matrix with first row [a, b, c, d], then:

det(M) = a * det(a_minor) - b * det(b_minor) + c * det(c_minor) - d * det(d_minor)

func Determinant(matrix [][]int) int {
      // your code here
}
*/

func Determinant(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0] // Base case for a 1x1 matrix
	}

	if len(matrix) == 2 {
		// Base case for a 2x2 matrix
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	determinant := 0

	for i := 0; i < len(matrix); i++ {
		// Create the minor matrix by excluding the first row and ith column
		minor := make([][]int, len(matrix)-1)
		for j := range minor {
			minor[j] = make([]int, len(matrix)-1)
		}

		for j := 1; j < len(matrix); j++ {
			for k := 0; k < len(matrix); k++ {
				if k != i {
					col := k
					if k > i {
						col = k - 1
					}
					minor[j-1][col] = matrix[j][k]
				}
			}
		}

		sign := 1
		if i%2 != 0 {
			sign = -1
		}

		determinant += sign * matrix[0][i] * Determinant(minor)
	}

	return determinant
}
