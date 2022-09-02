package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func writer(out chan<- int, durationWork int) {
	var isWrite bool = true
	rand.Seed(time.Now().UnixNano())

	go func() {
		<-time.After(time.Second * time.Duration(durationWork))
		isWrite = false
	}()

	for isWrite {
		out <- rand.Int()
	}
}

func reader(c <-chan int) {
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
}

var (
	durationWork int
)

func init() {
	flag.IntVar(&durationWork, "d", 1, "duration of work in seconds")
}

func main() {
	flag.Parse()

	c := make(chan int)

	reader(c)
	writer(c, durationWork)

	close(c)
}
