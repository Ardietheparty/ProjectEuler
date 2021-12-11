package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)
/*
10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

 */
//Prime Nums
var primes []int
func main() {
	timeStart := time.Now()
	//Init our first few primes
	primes = []int{2,3}
	NumOfPrimes := len(primes)
	n := 3
	LastPrime := 0
	//Only checking odds.
	for NumOfPrimes < 10002 {
		if IsPrime(n) {
			primes = append(primes,n)
			LastPrime = n
			NumOfPrimes++
		}
		n+=2
	}
	fmt.Println(LastPrime)
	fmt.Println(time.Since(timeStart))
}
func IsPrime(n int) bool {
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