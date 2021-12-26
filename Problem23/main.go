package main
/*
Non-abundant sums


A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

 */
import (
	"fmt"
	"time"
)
//SLOW should speed up
func main() {
	timeStart := time.Now()
	lm := 20161 //From wolfram alpha aylmao
	sum := 0
	var abuns []int
	for i := 12; i < lm+1; i++ {
		check := abun(i)
		if check {
			abuns = append(abuns,i)
		}
	}
	fmt.Println(time.Since(timeStart))
	timeStart = time.Now()
	for i := 1; i < lm+1; i++ {
		t := sumofabun(abuns,i)
		if t {
			sum += i
		}
	}
	fmt.Println(sum)

	fmt.Println(time.Since(timeStart))
}
func abun(n int) bool{
	test := divs(n)
	if test > n {
		return true
	}
	return false
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
		i++
	}
	return sum
}
func sumofabun(array []int, n int) bool {
	L := len(array)
	test := true
	s := 0
	for test {
		if n < array[s] {
			L = s
			test = false
		}
		if s == L - 2 {
			test = false
		}
		s++
	}
	for i := 0; i < L; i++ {
		for j := i; j < L; j++ {
			if n == array[i]+array[j] {
				return false
			}
		}
	}
	return true
}