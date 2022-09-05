package main

import "fmt"

func removeWithOrder(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	fmt.Println(removeWithOrder([]int{1, 2, 3, 4, 5}, 0))
}
