package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	personMap := make(map[string]string)
	var name string
	var address string

	fmt.Println("Enter name:")
	fmt.Scan(&name)
	fmt.Println("Enter address:")
	fmt.Scan(&address)
	
	personMap["name"] = name
	personMap["address"] = address

	jsonb, err := json.MarshalIndent(personMap, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonb))
}

