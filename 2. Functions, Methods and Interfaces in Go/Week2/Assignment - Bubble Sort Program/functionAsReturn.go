package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn is used to calculate the final displacement
func GenDisplaceFn(a, vO, sO float64) func(float64) float64 {

	timeFn := func(t float64) float64 {
		firstPart := 0.5 * a * math.Pow(t, 2)
		secondPart := vO * t
		thirdPart := sO
		// fmt.Println(firstPart)
		// fmt.Println(secondPart)
		// fmt.Println(thirdPart)
		return firstPart + secondPart + thirdPart
	}

	return timeFn
}

func main() {
	var time, acc, initVelocity, initDisplacement float64
	fmt.Print("Enter the Acceleration: ")
	fmt.Scan(&acc)
	// fmt.Println(acc)

	fmt.Print("Enter the Initial Velocity: ")
	fmt.Scan(&initVelocity)
	// fmt.Println(initVelocity)

	fmt.Print("Enter the Initial Displacement: ")
	fmt.Scan(&initDisplacement)
	// fmt.Println(initDisplacement)

	fmt.Print("Enter the Time: ")
	fmt.Scan(&time)
	// fmt.Println(time)

	fnTime := GenDisplaceFn(acc, initVelocity, initDisplacement)

	fmt.Println("Final displacement is: ", fnTime(time))

}
