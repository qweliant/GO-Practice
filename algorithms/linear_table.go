package main

import "fmt"

// Creating a one dimensional array or "Linear Table"
func main(){
	var scores = []int8{90, 70, 50, 80, 60, 85}

	// length variable
	var length = len(scores)

	// for loop printing elements in array
	for i := 0; i < length; i++{
		fmt.Printf("%d ", scores[i])
	}
}