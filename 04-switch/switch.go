package main

import (
	"fmt"
)

func main() {
	var c byte
	fmt.Printf("Do you want to subscribe me? (y/n)")
	fmt.Scanf("%c", &c)
	switch c {
	case 'y', 'Y':
		fmt.Println("Thank you!")
		fallthrough
	case 'n', 'N':
		fmt.Println("No problem")
	default:
		fmt.Println("Invalid input")
	}
}
