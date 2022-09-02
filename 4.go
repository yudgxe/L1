package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func Write(out chan<- int) {
	var isWrite bool = true

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, os.Interrupt)

	go func() {
		defer close(signalChanel)
		<-signalChanel
		isWrite = false
	}()

	rand.Seed(time.Now().UnixNano())

	for isWrite {
		out <- rand.Int()
	}

	return
}

func Read(c <-chan int, quantity int) {
	for i := 0; i < quantity; i++ {
		go func() {
			for v := range c {
				fmt.Println(v)
			}
		}()
	}
}

var (
	quantityReader int
)

func init() {
	flag.IntVar(&quantityReader, "q", 1, "quantity of reader")
}

func main() {
	flag.Parse()

	c := make(chan int)

	Read(c, quantityReader)
	Write(c)

	close(c)
}
