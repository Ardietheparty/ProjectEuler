package main

import (
	"fmt"
	"time"
)
/*
Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

 */
func main() {
	timeStart := time.Now()
	N := 2000000

	fmt.Println(time.Since(timeStart))
}
