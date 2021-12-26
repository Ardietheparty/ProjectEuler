package main

import (
	"fmt"
	"time"
)
/*
Longest Collatz sequence
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?
 */
var coll []int
func main() {
	timeStart := time.Now()
	l := 0
	temp := 0
	coll = make([]int, 10000000)
	fmt.Println(coll[1])
	//fmt.Println(collatz(13))
	for i := 1; i < 1000000; i++ {
		a := collatz(i)
		coll[i]=a
		if temp<coll[i]{
			temp = coll[i]
			l = i
		}
	}
	fmt.Println(temp,l)
	fmt.Println(time.Since(timeStart))
}
func odd(n int) int {
	return 3*n+1
}
func even(n int) int {
	return n / 2
}
//collatz Computes the number of steps required to get to 1
//it checks to see if you have found a collatz number before
func collatz(n int) int {
	i := n
	steps := 1
	for i != 1 {
		//fmt.Println(i)

		if i%2 == 0 {
			i = even(i)
		} else {
			i = odd(i)
		}
		steps++

		if i < len(coll) {
			if coll[i] != 0 {
				return steps+coll[i]
			}
		}


	}
	return steps
}