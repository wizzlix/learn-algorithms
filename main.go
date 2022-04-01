package main

/* Импорт пакета `fmt` и пакета `sorts`*/
import (
	"fmt"

	"./sorts"
)

func main() {
	/* Короткий способ объявить и инициализировать срез целых чисел. */
	a := []int{5, 7, 4, 5, 8, 3, 23, 53, 211, 2, 8, -4, 3}
	fmt.Println("Hello from wizzlix algorithms collection!")

	/* Вызов функции из другого пакета. */
	fmt.Println("Buble Sort :", sorts.BubbleSort(a))

}
