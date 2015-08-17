package sort

// Selection sort.
// The algorithm divides the input list into two parts:
// the sublist of items already sorted, which is built up from left to right at the front of the list,
// and the sublist of items remaining to be sorted that occupy the rest of the list.
// Initially, the sorted list is empty and the unsorted sublist is the entire input list.
// The algorithm proceeds by finding the smallest element in the unsorted list,
// exchanging it with the leftmost unsorted element, and moving the sublist boundaries one element to the right.
// Time complexity: W: O(n^2) B: O(n^2): A: O(n^2)
func SelectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := a[i]
		index := i

		for j := i; j < len(a); j++ {
			if a[j] < min {
				index = j
				min = a[j]
			}
		}
		a[i], a[index] = a[index], a[i]
	}
	return a
}
