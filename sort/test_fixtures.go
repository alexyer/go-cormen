package sort

import "math/rand"

const mainConst = 42

func init() {
	rand.Seed(mainConst)
}

func MakeRandArray() []int {
	return rand.Perm(mainConst)
}

func MakeSortedArray() []int {
	sortedArray := make([]int, mainConst)
	for i := 0; i < mainConst; i++ {
		sortedArray[i] = i
	}
	return sortedArray
}

func MakeReversedArray() []int {
	reversedArray := make([]int, mainConst)
	for i := 0; i < mainConst; i++ {
		reversedArray[mainConst-i-1] = i
	}
	return reversedArray
}
