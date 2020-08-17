package main

import (
	"fmt"
)

func main() {
	var appleNum float64
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&appleNum)
	println(appleNum)
	var y int32 = int32(appleNum)
	println(y)
}
