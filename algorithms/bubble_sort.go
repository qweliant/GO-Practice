package main

import "fmt"

func bubble(arr []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				var temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	// length variable
	var length = len(scores)

	bubble(scores, length)

	// for loop printing elements in array
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}
