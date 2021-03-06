package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)
/*
Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

 */
var primes []int
func main() {
	primes = []int{2}
	timeStart := time.Now()
	N := 2000000
	sum := 0
	for i := 3; i < N; i += 2 {
		if IsPrime(i) {
			primes = append(primes,i)
		}
	}
	for i:=0;i<len(primes);i++ {
		sum += primes[i]
	}
	fmt.Println(sum)

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