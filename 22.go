package main

import (
	"fmt"
	"math/big"
)

func div(x, y *big.Int) *big.Int {
	return big.NewInt(0).Div(x, y)
}

func mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

func add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

func sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}

func main() {
	bigX := new(big.Int)
	bigY := new(big.Int)

	bigX.SetString("100000000000000000000", 10)
	bigY.SetString("200000000000000000000", 10)

	fmt.Println(div(bigX, bigY).String())
	fmt.Println(mul(bigX, bigY).String())
	fmt.Println(add(bigX, bigY).String())
	fmt.Println(sub(bigX, bigY).String())
}
