package main

import (
	"fmt"
	"math/big"
	"time"
)
/*
1000-digit Fibonacci number


The Fibonacci sequence is defined by the recurrence relation:

    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.

Hence the first 12 terms will be:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

 */

func main() {
	timeStart := time.Now()
	lim := new(big.Int)
	lim.Exp(big.NewInt(10),big.NewInt(999),big.NewInt(0))

	fib := make([]*big.Int,3)
	fib[0] = big.NewInt(1)
	fib[2] = big.NewInt(1)
	fib[1] = big.NewInt(1)
	i := 0
	cnt := 2
	for fib[i].Cmp(lim) == -1 {
		i = (i+1)%3
		a := (i+1)%3
		b := (i+2)%3
		cnt++
		fib[i].Add(fib[a],fib[b])
		fmt.Println(cnt,fib[i],fib)
	}
	fmt.Println(cnt)
	fmt.Println(time.Since(timeStart))
}

