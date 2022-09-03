package main

import "fmt"

func flipString(input string) string {
	in := []rune(input)
	size := len(in)

	for i := 0; i < size/2; i++ {
		in[i], in[size-1-i] = in[size-1-i], in[i]
	}

	return string(in)
}

func main() {
	fmt.Println(flipString("главрыба"))
}
