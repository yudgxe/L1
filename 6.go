package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("quit")
				return
			default:
			}
		}
	}()

	quit <- true
	time.Sleep(10)
}
