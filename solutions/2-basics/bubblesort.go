package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 4, 2, 7, 5, 6, 10, 1, 9, 8}
	fmt.Println("Sorting array ", arr)
	sort(arr)
	fmt.Println("Done: ", arr)
}

func sort(arr []int) {
	done := false
	for !done {
		done = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				swap(arr, i, i+1)
				done = false
			}
		}
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}
