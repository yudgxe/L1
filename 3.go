package main

import (
	"fmt"
)

func sumSquares(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v * v
		}
	}()

	result := make(chan int)
	go func() {
		defer close(result)

		var sum int
		for v := range out {
			sum += v
		}

		result <- sum
	}()

	return result
}

func main() {
	fmt.Println(<-sumSquares([]int{2, 4, 6, 8, 10}))
}
