package sort

import (
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	sorted := MakeSortedArray()
	random := MakeRandArray()
	reversed := MakeReversedArray()

	InsertionSort(sorted)
	InsertionSort(random)
	InsertionSort(reversed)

	if !(sort.IntsAreSorted(sorted) && sort.IntsAreSorted(random) && sort.IntsAreSorted(reversed)) {
		t.Errorf("Wrong behavior:\n%v\n%v\n%v", sorted, random, reversed)
	}
}
