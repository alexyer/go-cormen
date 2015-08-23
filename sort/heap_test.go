package sort

import "testing"

func TestHeapSort(t *testing.T) {
	TestSort(t, HeapSort)
}
