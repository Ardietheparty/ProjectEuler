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
Maximum path sum I
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
 */
func main() {
	timeStart := time.Now()
	Data := nums("nums")
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
