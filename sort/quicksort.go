package sort

import (
	"math/rand"
	"time"
)

// Quicksort is divide and conquer algorithm.
// Steps:
//   1. Pick an element, called a pivot, from the array.
//   2. Reorder the array so that all elements with values less than the pivot come before the pivot,
//		while all elements with values greater than the pivot come after it.
//		After this partioning, the pivot is in its final position. This is called the partition operation.
//   3. Recursively apply the above steps to the sub-array of elements with smaller values to the sub-array of elements with greater values.
// Time complexity: W: O(n^2) B: O(n*log(n)) A: O(n*log(n))

type PartitionFn func([]int, int, int) int

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Straitforward implementation with Lomuto partition scheme
func QuickSort(a []int) []int {
	return Sort(a, 0, len(a)-1, Partition)
}

func Sort(a []int, lo, hi int, partition PartitionFn) []int {
	if lo < hi {
		q := partition(a, lo, hi)
		Sort(a, lo, q-1, partition)
		Sort(a, q+1, hi, partition)
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

// Randomized version of Quicksort.
// Choose pivot element randomly to obtain good expected performance over all inputs.
func RandomizedQuickSort(a []int) []int {
	return Sort(a, 0, len(a)-1, RandomizedPartition)
}

func RandomizedPartition(a []int, lo, hi int) int {
	i := rand.Intn(hi)
	a[hi], a[i] = a[hi], a[i]
	return Partition(a, lo, hi)
}

// The original partition scheme described by Hoare uses two indices that start at the ends of the array being partioned,
// then move toward each other, until they detect an inversion and swap inverted elements.
// When indices meet,the algorithm stops and returns the final index.
func HoareQuickSort(a []int) []int {
	return Sort(a, 0, len(a)-1, HoarePartition)
}

func HoarePartition(a []int, lo, hi int) int {
	x := a[lo]
	i := lo
	j := hi

	for {
		for a[j] > x {
			j--
		}

		for a[i] < x {
			i++
		}

		if i < j {
			a[i], a[j] = a[j], a[i]
		} else {
			return j
		}
	}
}
