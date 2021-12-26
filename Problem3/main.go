package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)
/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

 */
var primes []int
func main() {
	timeStart := time.Now()
	primes = []int{2,3}
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
func prime(n int) bool {
	pMax := primes[len(primes)-1]
	if n <= pMax {
		// If n is prime, it must be in the cache
		i := sort.SearchInts(primes, n)
		return n == primes[i]
	}
	max := int(math.Ceil(math.Sqrt(float64(n))))
	// Check if n is divisible by any of the cached primes
	for _, p := range primes {
		if p > max {
			return true
		}
		if n%p == 0 {
			return false
		}
	}
	// When you run out of cached primes, check if n is divisible by
	// any number 6*k+|-1 larger than the largest prime in the cache.
	for d := (pMax/6+1)*6 - 1; d <= max; d += 6 {
		if n%d == 0 || n%(d+2) == 0 {
			return false
		}
	}
	return true
}