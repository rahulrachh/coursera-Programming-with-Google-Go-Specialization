package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	names := make([]int, 3)
	scanner := bufio.NewScanner(os.Stdin)
	i := 0

	for {
		fmt.Print("Enter number: ")

		scanner.Scan()

		text := scanner.Text()
		textInt, _ := strconv.Atoi(text)

		if text != "X" {
			fmt.Println(text)
			if i > 2 {
				names = append(names, textInt)
				sort.Ints(names)
				fmt.Printf("%d\n", names)
			} else {
				j := sort.IntSlice(names).Search(0)
				names[j] = textInt
				sort.Ints(names)
				fmt.Printf("%d\n", names)
				i++
			}
		} else {
			break
		}
	}

	fmt.Println(names)
}
