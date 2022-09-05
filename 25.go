package main

import (
	"fmt"
	"time"
)

func sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	sleep(1)
	fmt.Println("hello")
}