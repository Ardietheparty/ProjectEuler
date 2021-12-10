package main

import (
	"fmt"
	"time"
)
/*
Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

 */

func main() {
	MAXIT := 1000000000000
	//generic way
	//Brute force
	timeStart := time.Now()
	var div []int
	for i := 2; i <= 20; i++ {
		div = append(div,i)
	}
	fmt.Println(div)
	for i := 2; i < MAXIT; i++ {
		test := mult(i,div)
		if test {
			fmt.Println(i)
			i = MAXIT+100
		}
	}
	fmt.Println(time.Since(timeStart))
	//Elgentish methode
	//Note that 2520 is divisable by 1,10 so we only need to check 11,20
	//If we only check multiples of 2520
	timeStart1 := time.Now()
	divsors := []int{11,12,13,14,15,16,17,18,19,20}
	for i:=2520; i < MAXIT; i+=2520 {
		test := mult(i,divsors)
		if test {
			fmt.Println(i)
			i = MAXIT+100
		}
	}
	fmt.Println(time.Since(timeStart1))
	//To scale
	//Best way
	// Find the smallest multiple of 1,2 then 1,2,3 then 1,2,3,4 etc
	timeStart2 := time.Now()
	biggestDivisor := 21
	hold := 2
	n := 0
	for i := 3; i < biggestDivisor+1; i++ {
		n = hold
		test := 0
		for test == 0 {
			if n%i == 0 {
				//fmt.Println(n)
				hold = n
				test =1
			} else {
				n += hold
			}

		}
	}
	fmt.Println(n)
	fmt.Println(time.Since(timeStart2))
}
func mult(i int, d []int) bool {
	for n := 0; n < len(d); n++ {
		if i%d[n] != 0 {
			return false
		}
	}
	return true

}