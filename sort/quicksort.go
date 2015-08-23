package sort

// Quicksort is divide and conquer algorithm.
// Steps:
//   1. Pick an element, called a pivot, from the array.
//   2. Reorder the array so that all elements with values less than the pivot come before the pivot,
//		while all elements with values greater than the pivot come after it.
//		After this partioning, the pivot is in its final position. This is called the partition operation.
//   3. Recursively apply the above steps to the sub-array of elements with smaller values to the sub-array of elements with greater values.
// Time complexity: W: O(n^2) B: O(n*log(n)) A: O(n*log(n))

// Straitforward implementation with Lomuto partition scheme
func QuickSort(a []int) []int {
	return Sort(a, 0, len(a)-1)
}

func Sort(a []int, lo, hi int) []int {
	if lo < hi {
		q := Partition(a, lo, hi)
		Sort(a, lo, q-1)
		Sort(a, q+1, hi)
	}
	return a
}

func Partition(a []int, lo, hi int) int {
	x := a[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if a[j] <= x {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
