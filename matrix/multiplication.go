package matrix

// Multiply 2 square matrices.
// Straightforward method.
// Time complexity: W: O(n^3) B: O(n^3) A: O(n^3)
func SquareMatrixMultiply(a, b [][]int) [][]int {
	n := len(a[0])
	res := allocMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res
}

// Allocate slice for matrix of n x n size.
func allocMatrix(n int) [][]int {
	matrix := make([][]int, n)
	vals := make([]int, n*n)
	for i := range matrix {
		matrix[i], vals = vals[:n], vals[n:]
	}
	return matrix
}
