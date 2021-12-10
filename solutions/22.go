package solutions

import (
	"fmt"
	"math/big"
)

// Task22 - 22.	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
func Task22() {
	// use big.Float for float numbers greater than ^18;
	a := new(big.Float)
	b := new(big.Float)

	a.SetString("6000000000000000000000")
	b.SetString("3000000000000000000000")

	sum := new(big.Float)
	sum.Add(a, b)
	fmt.Println("Sum: ", sum)

	sub := new(big.Float)
	sub.Sub(a, b)
	fmt.Println("Sub: ", sub)

	div := new(big.Float)
	div.Quo(a, b)
	fmt.Println("Div: ", div)

	mult := new(big.Float)
	mult.Mul(a, b)
	fmt.Println("Mult: ", mult)
}
