package algorithms

import "math"

func BinarySearch(a []int, n int) int {
	left := 0
	right := len(a) - 1

	for left <= right {
		mid := int(math.Round(float64((right-left)/2 + left)))

		if n == a[mid] {
			return mid
		} else if n < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
