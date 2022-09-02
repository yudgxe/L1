package main

import (
	"fmt"
)

// одна горутина считате все квадраты.
func squares(nums []int) <-chan int {
	out := make(chan int, len(nums))

	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v * v
		}
	}()

	return out
}

// len(nums) горутин считате квадраты.
func otherSquares(nums []int) <-chan int {
	out := make(chan int, len(nums))

	for i, v := range nums {
		go func(i int, v int) {
			out <- v * v

			if i == len(nums)-1 {
				close(out)
			}

		}(i, v)
	}

	return out
}

func main() {
	in := squares([]int{2, 4, 6, 8, 10})
	for v := range in {
		fmt.Println(v)
	}
}
