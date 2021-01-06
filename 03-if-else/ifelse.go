package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Printf("Enter a number : ")
	fmt.Scanf("%d", &num)
	fmt.Printf("%s", fizzBuzz(num))
}

func fizzBuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
		fallthrough
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}
