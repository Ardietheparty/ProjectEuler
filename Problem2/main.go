package main

import (
	"fmt"
	"time"
)

/*
Even Fibonacci numbers


Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

 */
func main() {
	t := time.Now()
	pre := []int{0,1,2}
	n := 3
	out := 0

	check := true
	for check == true {
		hld := pre[n-1]+pre[n-2]
		if hld>4000001{
			check = false
		} else {
			pre = append(pre, hld)
		}
		n++
	}
	for i := 0; i < len(pre); i++ {
		if pre[i]%2==0 {
			out += pre[i]
		}
	}
	fmt.Println(out)
	fmt.Println(time.Since(t))

}