package main

import (
	"fmt"
	"strconv"
)

func main() {
	var temp, sum int
	var err error
	var cmd string

	// Infinte loop
	for {
		fmt.Scanf("%s\n", &cmd)
		temp, err = strconv.Atoi(cmd)
		if err != nil {
			// Quit the loop
			break
		}
		sum = sum + temp
	}
	fmt.Printf("Sum of all values: %d", sum)
}
