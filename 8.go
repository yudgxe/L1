package main

import (
	"fmt"
	"os"
	"strconv"
)

func changeBit(value, position int64) int64 {
	return value ^ 1<<position
}

func main() {
	var v, p int64

	fmt.Fscan(os.Stdin, &v)
	fmt.Fscan(os.Stdin, &p)

	result := changeBit(v, p)

	fmt.Println(strconv.FormatInt(v, 2))
	fmt.Println(strconv.FormatInt(result, 2))
}
