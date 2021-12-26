package main

import (
	"fmt"
	"time"
)
/*
Coin Sums

In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:

    1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).

It is possible to make £2 in the following way:

    1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

How many different ways can £2 be made using any number of coins?

 */
func main() {
	timeStart := time.Now()
	coins := []int{1,2,5,10,20,50,100,200}
	sum := 200
	fmt.Println(NumberOfWays(sum,coins))

	fmt.Println(time.Since(timeStart))
}
func NumberOfWays(n int, coins []int) int {
	steps :=0
	ways := make([]int,n+1)
	ways[0]=1
	for i := 0; i < len(coins); i++ {
		for j := 0; j < len(ways); j++ {
			if coins[i] <= j {
				steps++
				ways[j] += ways[j-coins[i]]
			}
		}
	}
	fmt.Println(steps)
	return ways[n]
}