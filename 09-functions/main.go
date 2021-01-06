package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	x := sum(1, 2, 3, 4, 5)
	fmt.Println(x)
}

// Variadic functions takes any
// number of parameters of a given type
func sum(numbers ...int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

func readFile(name string) {
	var bytes [10]byte
	var err error

	f, err := os.Open(name)
	// Close function executes at the end of
	// this function even if the function crashes
	// in the middle
	defer f.Close()

	if err != nil {
		panic(err)
	}

	for {
		_, err := f.Read(bytes[:])
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", bytes)
	}
}
