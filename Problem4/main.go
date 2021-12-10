package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)
/*
Largest palindrome product


A palindromic number reads the same both ways. The largest palindrome made from the product
of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

 */
func main() {
	timeStart := time.Now()
	var all []int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			all = append(all,i*j)
		}
	}

	sort.Ints(all)
	I := len(all)
	for i := 0; i < I; i++ {
		num := all[I-i-1]
		check := palin(num)
		if check {
			fmt.Println(num)
			i=I*I
		}

	}
	fmt.Println(time.Since(timeStart))
}
func palin(i int) bool {
	s := strings.Split(strconv.Itoa(i),"")
	I := len(s)
	for n := 0; 2*n < I; n++ {
		if s[n]!=s[I-n-1] {
			return false
		}
	}
	return true
}
