package main

import (
	"fmt"
	"time"
)
/*
Special Pythagorean triplet


A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

 */
func main() {
	/*
	We solve by noting all primitive pythagorean triples can be found vy
	letting t>s>=1 and gcd(s,t)=1 where
	a = st, b=s^2-t^2 / 2, c = s^2+ t^2 /2
	Then all pythagorean triples can be found by
	a = 2st, b=s^2-t^2, c=s^2+t^2
	 */
	timeStart := time.Now()
	bnd := 1000
	fmt.Println(pythag(bnd))
	fmt.Println(time.Since(timeStart))
}
//Euclidean_Algo
func gcd(a int, b int) int {
	t := 0
	for b != 0 {
		t = b
		b = a % b
		a = t
	}
	return a
}
func A(s int, t int) int{
	return 2*s*t
}
func B(s int, t int) int{
	return s*s-t*t
}
func C(s int, t int) int {
	return s*s+t*t
}
func pythag(bound int) int {
	for t := 1; t < bound; t++ {
		for s := 1; s < t; s++ {
			a := A(t,s)
			b := B(t,s)
			c := C(t,s)
			if a+b+c == 1000 {
				fmt.Println(t,s)
				return a*b*c
			}
		}
	}

	return 0
}