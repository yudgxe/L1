package main

import "fmt"

type human struct {
	age int
}

func (h *human) getAge() int {
	return h.age
}

type action struct {
	human
}

func main() {
	a := &action{}
	fmt.Println(a.getAge())
}
