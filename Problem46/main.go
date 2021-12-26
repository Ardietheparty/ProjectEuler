package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)
/*
Golbach's other conjecture


It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2×12
15 = 7 + 2×22
21 = 3 + 2×32
25 = 7 + 2×32
27 = 19 + 2×22
33 = 31 + 2×12

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?

 */
var primes,  squares []int
func main() {
	bnd := 10000
	timeStart := time.Now()
	primes = []int{2,3,5,7}
	for i:=1;i<bnd;i++ {
		squares=append(squares,i*i)

	}
	for i := 9; i < bnd; i += 2 {
		if IsPrime(i) {
			primes = append(primes,i)
		}
	}
	fmt.Println(len(primes))
	check := true
	n := 11
	for check {
		n+=2
		check = gold(n)
	}
	fmt.Println(n)
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
	//primes = append(primes,n)
	return true
}

func GoldCheck(n int, p int, s int) bool {
	a := p + 2 * s
	if n == a {
		return true
	}
	return false
}
func gold(n int) bool {
	l  := len(primes)
	L := len(squares)
	if IsPrime(n) {
		return true
	}
	for i:=0;i<l;i++ {
		for s := 0; s < L; s++ {
			if GoldCheck(n, primes[i], squares[s]) {
				return true
			}
			if squares[s+1] >= n{
				s=L
			}
		}
	}
	return false
}