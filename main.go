package main

import (
	"fmt"

	"./sorts"
)

func main() {
	a := []int{5, 7, 4, 5, 8, 3, 23, 53, 211, 2, 8, -4, 3}
	fmt.Println("Hello from wizzlix algorithms collection!")

	fmt.Println("Buble Sort :", sorts.BubbleSort(a))

}
