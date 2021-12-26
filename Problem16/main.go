package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)
/*
Power digit sum

215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 21000?

 */

func main() {
	timeStart := time.Now()
	sum := 0
	a := new(big.Int).Exp(big.NewInt(2),big.NewInt(1000),nil).String()
	s := strings.Split(a, "")
	for i := 0; i < len(s); i++ {
		n,_:=strconv.Atoi(s[i])
		sum += n
	}

	fmt.Println(sum)

	fmt.Println(time.Since(timeStart))
}
