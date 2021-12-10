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
	timeStart := time.Now()
	divsors := []int{11,12,13,14,15,16,17,18,19,20}
	for i:=2520; i<= 100000000000; i+=2520 {
		test := mult(i,divsors)
		if test {
			fmt.Println(i)
			i=100000000000*2
		}
	}


	fmt.Println(time.Since(timeStart))
}
func mult(i int, d []int) bool {
	for n := 0; n < len(d); n++ {
		if i%d[n] != 0 {
			return false
		}
	}
	return true

}