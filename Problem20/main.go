package main

import (
	"fmt"
	"math/big"
	"time"
)
/*
Factorial digit sum


n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

 */
func main() {
	timeStart := time.Now()
	fac := new(big.Int)
	fac = big.NewInt(1)
	//fac.SetPrec(big.MaxPrec)
	for i := 1; i < 101; i++ {
		fac.Mul(fac,big.NewInt(int64(i)))
	}
	//q := new(big.Int)
	//q.DivMod(big.NewInt(10))
	sum := new(big.Int)
	sum = big.NewInt(0)
	a := new(big.Int)
	for fac.BitLen() != 0 {
		a.Mod(fac,big.NewInt(10))
		sum.Add(sum,a)
		fac.Div(fac,big.NewInt(10))
	}
	fmt.Println(sum)
	fmt.Println(fac)
	fmt.Println(time.Since(timeStart))
}
