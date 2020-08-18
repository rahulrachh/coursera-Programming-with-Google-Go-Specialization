package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	var addr string
	fmt.Printf("Enter your address: ")
	fmt.Scan(&addr)

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
