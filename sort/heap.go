package sort

// Heapsort can be thought of as an improved selection sort:
// like that algorithm, it divides its input into a sorted and an unsorted region,
// and it iteratively shrinks the unsorted region by extracting the largest element
// and moving that to the second region. The improvement consists of the use of a heap data structure
// rather than a linear-time search to find the maximum.
// Time complexity: W: O(n lgn) B: O(n lgn) A: (n lgn)
func HeapSort(a []int) []int {
	BuildMaxHeap(a)
	heapSize := len(a) - 1

	for i := len(a) - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapSize--
		MaxHeapify(a, 0, heapSize)
	}
	return a
}

// This function lets the value at a[i] float down in the max-heap so that the subtree
// at index i obeys the max-heap property.
func MaxHeapify(a []int, i, heapSize int) {
	left := i << 1
	right := i<<1 + 1
	var largest int

	if left <= heapSize && a[left] > a[i] {
		largest = left
	} else {
		largest = i
	}

	if right <= heapSize && a[right] > a[largest] {
		largest = right
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		MaxHeapify(a, largest, heapSize)
	}
}

// Goes through the remaining nodes of the tree and runs MaxHeapify on each one.
func BuildMaxHeap(a []int) {
	for i := ((len(a)) / 2) - 1; i >= 0; i-- {
		MaxHeapify(a, i, len(a)-1)
	}
}
