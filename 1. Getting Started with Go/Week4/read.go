package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	var FileName string
	fmt.Printf("Enter your FileName: ")
	fmt.Scan(&FileName)

	file, err2 := os.Open(FileName)
	if err2 != nil {
		fmt.Println(err2)
	}

	var newSlice []person

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new := strings.Split(scanner.Text(), " ")
		newSlice = append(newSlice, person{new[0], new[1]})

	}

	msg := "The slice of constructs is: "
	fmt.Printf("%s\n", msg)
	fmt.Printf("%v\n\n", newSlice)

	i := 0
	for range newSlice {
		fmt.Printf("First Name: %s, Last Name: %s\n", newSlice[i].fname, newSlice[i].lname)
		i++
	}

}
