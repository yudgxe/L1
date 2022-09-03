package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type counter struct {
	count int32
}

func (c *counter) increase() {
	atomic.AddInt32(&c.count, 1)
}

func main() {
	c := &counter{}
	go func() {
		c.increase()
	}()

	c.increase()
	time.Sleep(10)
	
	fmt.Println(c.count)
}
