package sort

import (
	"sort"
	"testing"
)

func TestCountingSort(t *testing.T) {
	sorted := MakeSortedArray()
	random := MakeRandArray()
	reversed := MakeReversedArray()
	b := make([]int, len(sorted))

	CountingSort(sorted, b, len(sorted))
	checkResult(b, t)

	CountingSort(random, b, len(sorted))
	checkResult(b, t)

	CountingSort(reversed, b, len(sorted))
	checkResult(b, t)
}

func checkResult(arr []int, t *testing.T) {
	if !sort.IntsAreSorted(arr) {
		t.Errorf("Wrong behavior:\n%v", arr)
	}
}
