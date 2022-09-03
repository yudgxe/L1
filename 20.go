package main

import (
	"fmt"
	"strings"
)

func flipString(input string) string {
	in := []rune(input)
	size := len(in)

	for i := 0; i < size/2; i++ {
		in[i], in[size-1-i] = in[size-1-i], in[i]
	}

	return string(in)
}

func flitpWords(input string) string {
	var sb strings.Builder
	words := strings.Fields(input)

	for _, word := range words {
		sb.WriteString(flipString(word))
		sb.WriteString(" ")
	}

	return sb.String()
}

func main() {
	fmt.Println(flitpWords("снег собака солнце"))
}
