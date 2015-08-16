package main

import (
	"fmt"

	"github.com/alexyer/go-cormen/sort"
)

func main() {
	a := []int{2, 1, 3}
	sort.InsertionSort(a)
	fmt.Println(a)
}
