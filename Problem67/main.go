package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
/*
Maximum path sum II
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem, as there are 299 altogether!
If you could check one trillion (1012) routes every second it would take over twenty billion years to check them all.
There is an efficient algorithm to solve it. ;o)
*/
func main() {
	timeStart := time.Now()
	Data := nums("nums.txt")
	//var holding []int
	for i := 0; i < len(Data); i++ {
		fmt.Println(Data[i],len(Data[i]))
	}
	//holding = append(holding,Data[0])
	var holder []int
	holding := Data[0]
	for i := 1; i < len(Data); i++ {
		holder = nil
		for j := 0; j < len(Data[i]); j++ {
			a:=holding[j]+Data[i][j]
			b:=holding[j+1]+Data[i][j]
			if a < b {
				holder = append(holder,b)
			} else {
				holder = append(holder,a)
			}
		}
		//fmt.Println(holding)
		holding = holder
	}
	fmt.Println(holding)

	fmt.Println(time.Since(timeStart))
}
func nums(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var hold [][]string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		temp := strings.Split(scan.Text()," ")
		hold = append(hold, temp)
	}
	var out [][]int
	//l := len(hold)
	//L := len(hold[l-1])
	/*
		for i := 0; i < l; i++ {
			c := make([]int,L)
			out = append(out,c)
		}

	*/

	//l := len(hold[len(hold)-1])
	for i := 0; i < len(hold); i++ {
		var temp []int
		for j := 0; j < len(hold[len(hold)-i-1]); j++ {
			n, _ := strconv.Atoi(hold[len(hold)-i-1][j])
			temp = append(temp,n)
		}
		out = append(out,temp)
	}
	return out
}
