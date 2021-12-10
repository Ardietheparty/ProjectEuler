package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()

	fmt.Println(time.Since(timeStart))
}
