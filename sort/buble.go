package sort

// Bubble sort.
// Simple algorithm that repeatedly steps through the list to be sorted,
// compares each pair of adjacent items and swaps them if they are in the wrong order.
// Time complexity: W: O(n^2) B: O(n) A: O(n^2)
func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
	return a
}
