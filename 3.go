package main

import (
	"fmt"
)

func sumSquares(nums []int) int {
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v * v
		}
	}()

	var sum int
	for v := range out {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(sumSquares([]int{2, 4, 6, 8, 10}))
}
