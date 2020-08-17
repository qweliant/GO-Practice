package main

import "fmt"

func shellSort(array []int32, length int) {

	var gaps = []int32{1391376, 463792, 198768, 86961, 33936,
		13776, 4592, 1968, 861, 336,
		112, 48, 21, 7, 3, 1}

	var r int32 = int32(length)
	var temp int32

	for k := 0; k < 16; k++ {
		for i := gaps[k]; i < r; i++ {
			temp = array[i]
			var j = i
			for j = i; j >= gaps[k] && array[j-gaps[k]] > temp; j -= gaps[k] {
				array[j] = array[j-gaps[k]]
			}
			array[j] = temp
		}
	}
}

func main() {

	var scores = []int32{10, 20, 30, 40, 590, 50, 60, 70, 80, 95, 210}
	var length = len(scores)

	shellSort(scores, length)

	// for loop printing elements in array
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}
