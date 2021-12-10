package main

import (
	"fmt"
	"math"
	"time"
)
/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

 */
func main() {
	timeStart := time.Now()
	var prims []int
	sum := 0
	N := 600851475143
	//N := 13195
	for i := 2; i*i < N; i++ {
		if N%i == 0 {
			if prime(i){
				prims = append(prims, i)
				sum = i
			}
		}
	}
	fmt.Println(prims)
	fmt.Println(sum)

	fmt.Println(time.Since(timeStart))
}
func prime(p int) bool{

	bnd := int(math.Ceil(math.Sqrt(float64(p))))+1
	for i := 2; i < bnd; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}
