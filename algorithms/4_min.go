package main

import "fmt"

/*
returns min value by searching thorugh index
*/
func min(array []int, length int, ) (int, int) {
	var index = 0
	var index2 = 1
	for i := 0; i < length; i++ {
		if array[index] > array[i] { // swap
			index2 = index
			index = i
		} else if array[index2] > array[i] {
			index2 = i
		}
	}
	return array[index], array[index2]
}

func main() {
	var scores = []int{90, 70, 50, 45, 100, 80, 60, 85}

	// length variable
	var length = len(scores)

	// nice to see this works
	var min, min2 = min(scores, length)
	fmt.Printf("Min Score = %d\n", min)
	fmt.Printf("2nd Min Score = %d\n", min2)

}
