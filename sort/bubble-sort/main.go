package main

import "fmt"

func bubbleSort(n []int) []int {
	for i := 0; i < len(n)-1; i++ {
		for j := 0; j < len(n)-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}

	return n
}

func main() {
	n := []int{2, 1, 5, 7, 9}
	fmt.Println(bubbleSort(n))
}
