package main

import (
	"fmt"

	"./algorithms"
	"./sorts"
)

func main() {
	a := []int{3, 4, 5, 67, 445, 345, -10, 67, 7, 65, 5, 54, 2, -23, 8, 9, 6754, 23, -2}
	sortedArray := sorts.BubbleSort(a)

	fmt.Println("")
	fmt.Println("BubbleSort:", sortedArray)
	fmt.Println("-------------------------")
	fmt.Println("BinarySearch index: ", algorithms.BinarySearch(sortedArray, 4))
	fmt.Println("")
}
