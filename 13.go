package main

import "fmt"

func swap(a, b *int) {
	*a = *a - *b // a = 9 - 4 = 5
	*b = *a + *b // b = 5 + 4 = 9
	*a = *b - *a // a = 9 - 5
}

func main() {
	a := 10
	b := 5

	swap(&a, &b)

	fmt.Println(a)
	fmt.Println(b)
}
