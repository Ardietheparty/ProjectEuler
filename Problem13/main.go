package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)
/*
Large Sum
Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
 */
func main() {
	timeStart := time.Now()
	sum := new(big.Int)
	NUM := numbers("numbers")
	for i := 0; i < len(NUM); i++ {
		sum.Add(sum,NUM[i])
	}
	fmt.Println(sum)
	fmt.Println(time.Since(timeStart))
}
func numbers(filename string) []*big.Int {
	Nums := make([]*big.Int,0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error getting file")
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		hold := scan.Text()
		n := new(big.Int)
		n, e := n.SetString(hold,10)
		if !e {
			fmt.Println("SetString Error")
		}
		Nums = append(Nums, n)
	}
	return Nums
}
