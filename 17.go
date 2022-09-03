package main

import (
	"fmt"
	"math"
)

func binarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		m := int(math.Floor((float64(low+high) / 2)))
		mid := arr[m]

		if mid < key {
			low = m + 1
		} else if mid > key {
			high = m - 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7}, 2))
}
