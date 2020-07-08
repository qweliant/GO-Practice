package main

import "fmt"


/*
returns min value by searching thorugh index
*/
func min(array []int, length int) int {
	var index = 0
	for i := 1; i < length; i++ {
		if array[index] > array[i] { // swap
			index = i
		}
	}

	return array[index]
}

func main() {
	var scores = []int{90, 70, 50, 45, 100, 80, 60, 85}

	// length variable
	var length = len(scores)

	// nice to see this works
	var min = min(scores, length)
	fmt.Printf("Min Score = %d\n", min)
}
