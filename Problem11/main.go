package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)
/*
Largest product in a grid
In the 20×20 grid below, four numbers along a diagonal line have been marked in red.
The product of these numbers is 26 × 63 × 78 × 14 = 1788696.

What is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20×20 grid?
Note grid is in grid.
 */

func main() {
	timeStart := time.Now()
	grid := getGrid("grid")
	H := len(grid)
	L := len(grid[0])


	a := CheckDown(grid, H, L)
	b := CheckRight(grid, H, L)
	c := CheckDiagUp(grid, H, L)
	d := CheckDiagDown(grid, H, L)
	out := []int{a,b,c,d}
	sort.Ints(out)
	fmt.Println(out)
	fmt.Println(out[len(out)-1])
	//fmt.Println(grid[1][2])
	fmt.Println(time.Since(timeStart))
}
//getGrid Constructs our grid from our grid file
//Goes from string to int
func getGrid(filename string) [][]int {
	var line []string
	var hold []int
	var grid [][]int
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Couldn't Find Grid :(")
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line = append(line, scan.Text())
	}
	for i := 0; i < len(line); i++ {
		temp := strings.Split(line[i]," ")
		hold = nil
		for j := 0; j < len(temp); j++ {
			num, _ := strconv.Atoi(temp[j])
			hold = append(hold,num)
		}
		grid = append(grid,hold)
	}
	return grid
}
func CheckRight(grid [][]int, H int, L int) int {
	prod := 0
	for i := 0; i < H; i++ {
		for j := 0; j < L-4; j++ {
			tes := grid[i][j]*grid[i][j+1]*grid[i][j+2]*grid[i][j+3]
			if prod < tes {
				prod = tes
			}
		}
	}
	return prod
}
func CheckDown(grid [][]int, H int, L int) int {
	prod := 0
	for i := 0; i < H-4; i++ {
		for j := 0; j < L; j++ {
			tes := grid[i][j]*grid[i+1][j]*grid[i+2][j]*grid[i+3][j]
			if prod < tes {
				prod = tes
			}
		}
	}
	return prod
}
func CheckDiagDown(grid [][]int, H int, L int) int {
	prod := 0
	for i := 0; i < H-4; i++ {
		for j := 0; j < L-4; j++ {
			tes := grid[i][j]*grid[i+1][j+1]*grid[i+2][j+2]*grid[i+3][j+3]
			if prod < tes {
				prod = tes
			}
		}
	}
	return prod
}
func CheckDiagUp(grid [][]int, H int, L int) int {
	prod := 0
	for i := 0; i < H-4; i++ {
		for j := 0; j < L-4; j++ {
			l := L-1
			tes := grid[i][l-j]*grid[i+1][l-j-1]*grid[i+2][l-j-2]*grid[i+3][l-j-3]
			if prod < tes {
				prod = tes
			}
		}
	}
	return prod
}

