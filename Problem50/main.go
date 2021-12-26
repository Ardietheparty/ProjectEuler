package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)
/*
Consecutive prime sum


The prime 41, can be written as the sum of six consecutive primes:
41 = 2 + 3 + 5 + 7 + 11 + 13

This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?

 */
var primes []int
func main() {
	timeStart := time.Now()
	bnd := 100
	primes = []int{2}
	for i := 3; i < bnd; i++ {
		if IsPrime(i) {
			primes = append(primes,i)
		}
	}
	primeSum := make([]int,len(primes)+1)
	primeSum[0]=0
	for i := 0; i < len(primes); i++ {
		primeSum[i+1]=primeSum[i]+primes[i]
	}
	fmt.Println(primeSum)

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