package main

import "fmt"

func main() {
	// Declare a variable to store a slice of integers
	var mySlice []int

	// Append some numbers to the slice
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)

	// Print out the contents of the slice
	fmt.Println("My Slice:", mySlice)
}
