package main

import "fmt"

/*
The code here is taken from problems 1,  6, 7, 8, 11,
*/
// Creating a one dimensional array or "Linear Table"
func append(){}

func main() {
	var scores = []int8{90, 70, 50, 80, 60, 85}

	// length variable

	var length = len(scores)

	// append value to array, we need a temp array larger than scores
	var tempArray = make([]int8, length+1)

	// for loop printing elements in array
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}

	// copy values from scores into temp array
	for i := 0; i < length; i++ {
		tempArray[i] = scores[i]
	}
	fmt.Printf("\n")
	// assign value to end
	tempArray[length] = 75

	// change pointer ref to temp
	scores = tempArray
	for i := 0; i < length+1; i++ {
		fmt.Printf("%d ", scores[i])
	}

	fmt.Printf("\nSuccess!\n")

}
