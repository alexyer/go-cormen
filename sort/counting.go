package sort

// Counting sort operates by counting the number of objects that have each distinct key value,
// and using arithmetic on those counts to determine the positions of each key value in the output sequence.
// Its running time is linear in the number of items and the difference between the maximum and minimum key values,
// so it is only suitable for direct use in situations where the variation in keys is not significantly greater than
// the number of items. However, it is often used as a subroutine in radix sort, that can handle larger keys more efficiently.
// Time Complexity: W: O(n), B: O(n), A: O(n)
func CountingSort(a, b []int, k int) {
	c := make([]int, k)

	for i := 0; i < len(a); i++ {
		c[a[i]] += 1
	}

	for i := 1; i < k; i++ {
		c[i] += c[i-1]
	}

	for i := len(a) - 1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]] = c[a[i]] - 1
	}
}
