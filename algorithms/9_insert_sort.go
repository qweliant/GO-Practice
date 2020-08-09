package main
import "fmt"

/*
start with first element in list
if the next is less, place it to the left of it, else move on
*/

func main() {
	var scores = []int{80, 70, 60, 50, 95}
	var length = len(scores)

	sort(scores, length)

	// for loop printing elements in array
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func sort(arr []int, length int) {	
	
	// iterate over the array
	// if arr[i] is greater than arr[i+1], arr[i] = arr[i+1], arr[i+1] = arr[i]
	for i := 0; i < length-1; i++ {
		
		var temp = arr[i] //unsorted new element
		var pos = i //positio nvariable
		for j:= pos - 1; j >= 0; j--{ 
			// check if it less than all values in the array this position
			if temp < arr[j] {
				arr[j + 1] = arr[j]
				pos--
			}
		}
		arr[pos] = temp
	}
}


