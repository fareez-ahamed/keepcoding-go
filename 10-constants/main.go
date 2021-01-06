package main

import "fmt"

// DayOfWeek represents a day
type DayOfWeek int

// Day of week constants
const (
	Sunday DayOfWeek = 1 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Friday)
}
