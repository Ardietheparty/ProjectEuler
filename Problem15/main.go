package main

import (
	"fmt"
	"time"
)
/*
Lattice paths
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
How many such routes are there through a 20×20 grid?
 */
func main() {
	//Solve using binomial coefficient  is just 2n choose n
	timeStart := time.Now()
	n := 20
	prod := 1
	for i := 1; i < n+1; i++ {
		prod *= 2*n + 1 - i
		prod /= i
	}
	fmt.Println(prod)
	fmt.Println(time.Since(timeStart))
}
