/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the
characters ‘i’, ‘a’, and ‘n’. The program should print “Found!”
if the entered string starts with the character ‘i’, ends with
the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter
if the characters are upper-case or lower-case.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter the string with i, a and n: ")
	var y string
	fmt.Scan(&y)
	x := strings.ToLower(y)
	// var pref string = 'a'
	if strings.HasPrefix(x, "i") {
		if strings.HasSuffix(x, "n") {
			if strings.Contains(x, "a") {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}
		} else {
			fmt.Println("Not found!")
		}
	} else {
		fmt.Println("Not found!")
	}
}
