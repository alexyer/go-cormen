package sort

import "testing"

func TestQuickSort(t *testing.T) {
	TestSort(t, QuickSort)
}

func TestRandomizedQuickSort(t *testing.T) {
	TestSort(t, RandomizedQuickSort)
}

func TestHoareQuickSort(t *testing.T) {
	TestSort(t, HoareQuickSort)
}
