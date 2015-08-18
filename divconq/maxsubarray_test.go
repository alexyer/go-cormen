package divconq

import "testing"

type TestCase struct {
	array     []int
	crossLow  int
	crossHigh int
	crossSum  int
}

var testCases = []TestCase{
	TestCase{[]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}, 7, 10, 43},
	TestCase{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 3, 6, 6},
}

func TestMaxSubarray(t *testing.T) {
	for _, testCase := range testCases {
		crossLow, crossHigh, crossSum := MaxSubarray(testCase.array)
		if crossLow != testCase.crossLow && crossHigh != testCase.crossHigh && crossSum != testCase.crossSum {
			t.Errorf("Wrong result: %d, %d, %d\nExpected: %d, %d, %d", crossLow, crossHigh, crossSum, testCase.crossLow, testCase.crossHigh, testCase.crossSum)
		}
	}
}
