package main

import (
	"fmt"
	"math/rand"
	"time"
)

func intersection(farr []int, sarr []int) []int {
	m := make(map[int]bool)
	var result []int

	for _, v := range farr {
		m[v] = true
	}

	for _, v := range sarr {
		if s, ok := m[v]; ok && s {
			result = append(result, v)
		}
	}

	return result
}

func randIntArr(size, min, max int) []int {
	rand.Seed(time.Now().UnixNano())

	result := make([]int, size)
	for i := range result {
		result[i] = rand.Intn(max-min) + min
	}

	return result
}

func main() {
	farr := randIntArr(10, 0, 10)
	sarr := []int{1, 4, 5, 7}

	fmt.Println(farr)
	fmt.Println(sarr)

	r := intersection(farr, sarr)
	fmt.Println(r)
}