package sort

func InsertionSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]

		var i int
		for i = j - 1; i >= 0 && a[i] > key; i-- {
			a[i+1] = a[i]
		}
		a[i+1] = key
	}
	return a
}
