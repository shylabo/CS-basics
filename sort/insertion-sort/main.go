package main

import "fmt"

func insertionSort(n []int) []int {
	for i := 1; i < len(n); i++ {
		for j := 0; j < i; j++ {
			if n[i-j-1] > n[i-j] {
				n[i-j-1], n[i-j] = n[i-j], n[i-j-1]
			} else {
				break
			}
		}
	}
	return n
}

func main() {
	n := []int{3, 1, 4, 6, 7, 2, 9, 8}
	fmt.Println(insertionSort(n))
}
