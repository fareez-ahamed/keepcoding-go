package main

import "fmt"

func main() {
	var listA []int
	listA = append(listA, 1, 2)
	listB := make([]int, len(listA))
	copy(listB, listA)
	listB[0] = 3
	printSlice(listA)
	printSlice(listB)
}
func printSlice(slice []int) {
	fmt.Printf("%v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))
}
