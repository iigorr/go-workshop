package main

import (
	"fmt"
)

//TODO

func main() {
	arr := []int{3, 4, 2, 7, 5, 6, 10, 1, 9, 8}
	fmt.Println("Sorting array ", arr)
	sort(arr)
	fmt.Println("Done: ", arr)
}
