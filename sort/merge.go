package sort

// Merge sort.
// Divide and conquer algorithm.
// Have 2 steps:
//	1. Divide unsorted list into n sublists, each containing 1 element (a list of 1 element is considered sorted).
//	2. Repetedly merge sublists to produce new sorted sublists until there is only 1 sorted sublist remaining.
// Time complexity: W: O(n log n) B: O(n log n) A: O(n log n)
func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	// Recursively split runs into two halves until run size == 1,
	// then merge them and return back up call chain.
	middle := len(array) / 2
	a := MergeSort(array[:middle])
	b := MergeSort(array[middle:])
	return merge(a, b)
}

func merge(a, b []int) []int {
	r := make([]int, len(a)+len(b))
	i := 0
	j := 0

	// Copy smallest element of 2 sublist and move pointer
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	// Copy left elements
	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}
