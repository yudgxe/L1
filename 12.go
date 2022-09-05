package main

import "fmt"

type set struct {
	set map[string]struct{}
}

func new() *set {
	return &set{
		set: make(map[string]struct{}),
	}
}

func (s *set) Add(str string) {
	s.set[str] = struct{}{}
}

func main() {
	set := new()

	set.Add("cat")
	set.Add("cat")
	set.Add("dog")
	set.Add("cat")
	set.Add("set")

	fmt.Println(set.set)
}
