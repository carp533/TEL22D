package main

import (
	"fmt"

	"github.com/carp533/TEL22D/filter"
)

func swap(a []int, i, k int) {
	//h := a[i]
	//a[i] = a[k]
	//a[k] = h
	a[i], a[k] = a[k], a[i]
}

func bubbleSort(a []int) {
	var (
		n      = len(a)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i+1], a[i] = a[i], a[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	//	hi := len(a) - 1
	//	for k := hi; k > 0; k-- {
	//		for i := 0; i < k; i++ {
	//			if a[i] > a[i+1] {
	//				swap(a, i, i+1)
	//			}
	//		}
	//	}
}

func main() {

	l1 := []int{10, 3, 17, 5, 23, 12}
	bubbleSort(l1)
	fmt.Println(l1)
	lesser := filter.FilterLess(l1, 11)
	fmt.Println(lesser)
}
