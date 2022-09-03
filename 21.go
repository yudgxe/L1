package main

import "fmt"

// базовый интерфейс
type iTarget interface {
	getRequest() string
}

//реализация
type stringTarget struct{}

func (st *stringTarget) getRequest() string { return "string" }

//базовый класс используюший ITarget
type Client struct {
	iTarget
}

func (c *Client) Read() {
	fmt.Println(c.getRequest())
}

// адаптируемый класс
type intTarget struct{}

func (it *intTarget) getRequest() int {
	return 1 << 1
}

// Адаптер
type adapter struct {
	iTarget
	it intTarget
}

func (a *adapter) getRequest() string {
	return string(a.it.getRequest())
}

func main() {
	c := &Client{
		&adapter{},
	}

	c.Read()
}
