package main

import "fmt"

/*
Multiples of 3 or 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.

 */
func main() {
	var sum int
	for i := 0; i < 10001; i++ {
		if i%3 == 0 {
			sum+=i
		} else if i%5 == 0 {
			sum+=i
		}
	}
	fmt.Println(sum)
}
