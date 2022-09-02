package main

import (
	"fmt"
)

func dozensSort(arr []float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, v := range arr {
		key := int(v) / 10 * 10
		m[key] = append(m[int(key)], v)
	}
	return m
}

func main() {
	temepratura := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := dozensSort(temepratura)

	fmt.Println(m)
}
