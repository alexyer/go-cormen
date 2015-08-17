package sort

import (
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	sorted := MakeSortedArray()
	random := MakeRandArray()
	reversed := MakeReversedArray()

	SelectionSort(sorted)
	SelectionSort(random)
	SelectionSort(reversed)

	if !(sort.IntsAreSorted(sorted) && sort.IntsAreSorted(random) && sort.IntsAreSorted(reversed)) {
		t.Errorf("Wrong behavior:\n%v\n%v\n%v", sorted, random, reversed)
	}
}
