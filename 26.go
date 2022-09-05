package main

import (
	"fmt"
	"strings"
)

func count(in string) bool {
	low := strings.ToLower(in)
	m := make(map[rune]int)

	for _, v := range low {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] += 1
	}

	return true
}

func main() {
	fmt.Println(count("abcd"))
}
