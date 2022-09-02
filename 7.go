package main

import (
	"fmt"
	"sync"
)

type cmap[F comparable, T any] struct {
	mutex sync.RWMutex
	m     map[F]T
}

func (c *cmap[F, T]) write(k F, v T) {
	if len(c.m) == 0 {
		c.m = make(map[F]T)
	}

	c.mutex.Lock()
	c.m[k] = v
	c.mutex.Unlock()
}

func main() {
	c := &cmap[int, int]{}
	c.write(1, 100)

	fmt.Println(c.m)
}
