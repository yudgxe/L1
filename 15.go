package main

import "fmt"

var justString string

/*
var justString string глобальная?
func someFunc() {
  v := createHugeString(1 << 10) хардкод значений? 1 << 10 размер? сид?
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

//если createHugeString принемает любой параметр, но не размер
func createHugeString(seed int) string {
	return ""
}

func someFunc(seed, size int) string {
	return createHugeString(seed)[:size]
}

func main() {
	fmt.Println(someFunc(1024, 100))
}
