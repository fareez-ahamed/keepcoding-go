package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	wheelsMap := map[string]int{
		"car":     4,
		"auto":    3,
		"bicycle": 2,
	}

	data, _ := json.Marshal(wheelsMap)

	fmt.Printf("%s", data)
}
