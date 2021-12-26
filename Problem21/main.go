package main

import (
	"fmt"
	"time"
)
/*
Amicable numbers


Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

 */
var steps int
func main() {
	steps = 0
	timeStart := time.Now()
	N := 10000
	div := []int{1}
	for i := 2; i < N; i++ {
		div = append(div,divs(i))
	}
	/*
	for i := 0; i < len(div); i++ {
		fmt.Println(i+1,div[i])
	}
	temp := 0
	for i := 0; i < len(div); i++ {
		if temp < div[i] {
			temp = div[i]
		}
	}
	 */
	sum := 0
	for i := 1; i < N; i++ {
		a:=i
		da:=div[a-1]
		b:=da
		db:=0
		//63272
		if b-1<len(div) {
			db = div[b-1]
		} else {
			db = divs(b)
		}
		if db == a && da != a{
			sum += a+b
			//fmt.Println(a,da,b,db)
		}

	}
	/*
	for i := 0; i < len(div); i++ {
		if div[i] != 0 {
			for j := 0; j < len(div); j++ {
				if i != j {
					if div[i] == div[j] {
						sum += i+j+2
						fmt.Println(i+1,j+1)
						//fmt.Println(div[i],div[j])
						pairs++
						div[i]=0
						div[j]=0
						j=len(div)
					}
				}

			}
		}
	}

	 */
	fmt.Println(sum/2)
	fmt.Println(time.Since(timeStart))
}
func divs(n int) int {
	i:=2
	sum := 1
	j:=n
	for i < j {
		if n%i == 0 {
			h:=n/i
			if h*h==n {
				sum-=h
			}
			sum += i + h
			j = h
		}
		steps++
		i++
	}
	return sum
}
