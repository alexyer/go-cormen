package matrix

import "testing"

type TestCase struct {
	a      [][]int
	b      [][]int
	result [][]int
}

var testCases = []TestCase{
	TestCase{
		[][]int{[]int{1, 2}, []int{3, 4}},
		[][]int{[]int{5, 6}, []int{7, 8}},
		[][]int{[]int{19, 22}, []int{43, 50}},
	},
	TestCase{
		[][]int{[]int{42, -73}, []int{-73, 42}},
		[][]int{[]int{1, 2}, []int{3, 4}},
		[][]int{[]int{-177, -208}, []int{53, 22}},
	},
}

func TestSquareMatrixMultiply(t *testing.T) {
	for _, testCase := range testCases {
		result := SquareMatrixMultiply(testCase.a, testCase.b)
		for i := 0; i < len(result[0]); i++ {
			for j := 0; j < len(result[0]); j++ {
				if result[i][j] != testCase.result[i][j] {
					t.Errorf("Wrong result: %v\nExpected: %v", result, testCase.result)
					return
				}
			}
		}
	}
}
