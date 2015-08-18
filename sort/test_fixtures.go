package sort

import (
	"math/rand"
	"sort"
	"testing"
)

const mainConst = 42

type SortFunction func([]int) []int

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

func TestSort(t *testing.T, sortFn SortFunction) {
	sorted := MakeSortedArray()
	random := MakeRandArray()
	reversed := MakeReversedArray()

	sorted = sortFn(sorted)
	random = sortFn(random)
	reversed = sortFn(reversed)

	if !(sort.IntsAreSorted(sorted) && sort.IntsAreSorted(random) && sort.IntsAreSorted(reversed)) {
		t.Errorf("Wrong behavior:\n%v\n%v\n%v", sorted, random, reversed)
	}
}
