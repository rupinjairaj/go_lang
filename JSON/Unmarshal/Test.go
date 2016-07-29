package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
    {"Name": "Playptus", "Order": "Monotremata"},
    {"Name": "Quoll", "Order": "Dasyuromorphia"}
    ]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		panic(err)
	}
	// fmt.Println(jsonBlob)
	fmt.Println(animals)
}
