package main

import (
	"fmt"
)

func selectionSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		min := i
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[min] {
				min = j
			}
		}

		n[i], n[min] = n[min], n[i]
	}

	return n
}

func main() {
	n := []int{3, 6, 2, 5, 8, 9, 1}
	// fmt.Println(selectionSort(n))
	fmt.Println(selectionSort(n))
}
