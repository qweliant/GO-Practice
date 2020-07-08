package main

import "fmt"

// sort array in ascending order and return max val
func max(array []int, length int) int {
	for i := 0; i < length-1; i++ {
		if array[i] > array[i+1] { // swap
			var temp = array[i]
			array[i] = array[i+1]
			array[i+1] = temp
		}
	}
	var maxVal = array[length-1]
	return maxVal
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	// length variable
	var length = len(scores)
	var maxVal = max(scores, length)
	fmt.Printf("High Score = %d\n", maxVal)
}
