package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	fmt.Printf("Enter your name: ")
	readsname := bufio.NewReader(os.Stdin)
	name, err1 := readsname.ReadString('\n')
	if err1 != nil {
		println("Errror")
	}

	var addr string
	fmt.Printf("Enter your address: ")
	readsaddr := bufio.NewReader(os.Stdin)
	addr, err2 := readsaddr.ReadString('\n')

	if err2 != nil {
		println("Errror")
	}

	var idMap map[string]string
	idMap = make(map[string]string)
	idMap["name"] = name
	idMap["address"] = addr

	// for key, val := range idMap {
	// 	fmt.Println(key, val)
	// }

	barr, _ := json.Marshal(idMap)
	// fmt.Println(string(barr))
	os.Stdout.Write(barr)

}
