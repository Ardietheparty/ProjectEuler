package main

import (
	"fmt"
	"time"
)
/*
Lexicographic permutations


A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

 */
var cnt int
var st []string
func main() {
	timeStart := time.Now()
	//perm:=0
	s := "0123456789"
	//s := "012345"
	cnt = 0
	dig := []rune(s)
	perms(dig,0,len(dig)-1)
	//sort.Strings(st)
	fmt.Println(st[999999])

	fmt.Println(time.Since(timeStart))
}
func perms(s []rune, left, right int)  {
	if cnt == 1000000 {
		//return
	}
	if left == right {
		//fmt.Println(cnt,string(s))
		st=append(st,string(s))
		cnt++
	} else {
		for i := left; i <= right; i++ {
			s[left],s[i] = s[i],s[left]
			perms(s,left+1,right)
			s[left], s[i] = s[i],s[left]
		}
	}

}
