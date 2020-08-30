package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap function to swap in BubbleSort
func Swap(numSlice []int, innerIndex int) {
	numSlice[innerIndex], numSlice[innerIndex-1] = numSlice[innerIndex-1], numSlice[innerIndex]
}

// BubbleSort function to efficiently implemented using slices
func BubbleSort(numSlice []int) {

	for index := len(numSlice); index > 0; index-- {

		for innerIndex := 1; innerIndex < index; innerIndex++ {

			if numSlice[innerIndex-1] > numSlice[innerIndex] {

				Swap(numSlice, innerIndex)

			}
		}
	}

}

// MakingSlice function to store the integers in integers slice
func MakingSlice(input string) []int {
	var num []int
	for _, numbers := range strings.Fields(input) {
		numbersInIntegers, err := strconv.Atoi(numbers)
		if err == nil {
			num = append(num, numbersInIntegers)
		}
	}
	return num
}

func main() {
	fmt.Print("Enter number: ")
	scanner := bufio.NewScanner(os.Stdin)
	var a []int

	for scanner.Scan() {
		a = MakingSlice(scanner.Text())
		fmt.Println("This is the slice BEFORE sorting: ", a)
		break
	}

	BubbleSort(a)
	fmt.Println("This is the slice AFTER sorintg: ", a)

}
