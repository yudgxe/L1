package main

import (
	"fmt"
	"reflect"
)

// int, string, string, channel
func doType(variable interface{}) {
	v := reflect.ValueOf(variable)
	switch v.Kind() {
	case reflect.Int:
		fmt.Printf("int: %v\n", v.Int())
	case reflect.Bool:
		fmt.Printf("bool: %v\n", v.Bool())
	case reflect.String:
		fmt.Printf("string: %v\n", v.String())
	case reflect.Chan:
		fmt.Printf("chan: %v\n", v.Interface())
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	doType(1)
	doType("a")
	doType(false)
	doType(make(chan int))
}
