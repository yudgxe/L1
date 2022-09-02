package main

import "fmt"

func genrate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func main() {
	gen := genrate(2, 4, 6, 8, 10)
	sq := square(gen)

	for v := range sq {
		fmt.Println(v)
	}
}
