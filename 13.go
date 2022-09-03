package main

import "fmt"

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func main() {
	a := 10
	b := 5

	swap(&a, &b)
	
	fmt.Println(a)
	fmt.Println(b)
}
