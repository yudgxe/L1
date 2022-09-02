package main

import (
	"fmt"
)

func printChan(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

// одна горутина считате все квадраты.
func squares(nums []int) {
	out := make(chan int, len(nums))

	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v * v
		}
	}()

	printChan(out)
}

// len(nums) горутин считате квадраты.
func otherSquares(nums []int) {
	out := make(chan int, len(nums))

	for i, v := range nums {
		go func(i int, v int) {
			out <- v * v

			if i == len(nums)-1 {
				close(out)
			}

		}(i, v)
	}

	printChan(out)
}

func main() {
	otherSquares([]int{2, 4, 6, 8, 10})
}
