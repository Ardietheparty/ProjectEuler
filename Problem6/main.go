package main

import (
	"fmt"
	"time"
)
/*
Sum square difference


The sum of the squares of the first ten natural numbers is,
$$1^2 + 2^2 + ... + 10^2 = 385$$

The square of the sum of the first ten natural numbers is,
$$(1 + 2 + ... + 10)^2 = 55^2 = 3025$$

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is $3025 - 385 = 2640$.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

 */
func main() {
	n := 10000
	timeStart := time.Now()

	//Boring way by brute forces
	fmt.Println(SquareOfSum(n)-SumOfSquares(n))
	fmt.Println(time.Since(timeStart))

	//Gauss way
	timeStart1 := time.Now()
	sumSquare := (n*(n+1)/2)*(n*(n+1)/2)
	squaredSum := (n*(n+1)*(2*n+1))/6
	fmt.Println(sumSquare-squaredSum)
	fmt.Println(time.Since(timeStart1))
}
func SumOfSquares(i int) int{
	ret :=0
	for n := 1; n < i+1; n++ {
		ret+=n*n
	}
	return ret
}
func SquareOfSum(i int) int {
	hold := 0
	for n := 1; n < i+1; n++ {
		hold += n
	}
	return hold*hold
}
