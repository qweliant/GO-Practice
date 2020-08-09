package main

import "fmt"

/*
The code here is taken from problems 1,  6, 7, 8, 10, 11
*/
// Creating a one dimensional array or "Linear Table"

// appends value to array, we need a temp array larger than scores
// requires an array and the value to append to that array
func append(array []int8, val int8) []int8 {
	var length = len(array)
	var tempArray = make([]int8, length+1)

	// copy values from scores into temp array
	for i := 0; i < length; i++ {
		tempArray[i] = array[i]
	}
	fmt.Printf("\n")

	// assign value to end
	tempArray[length] = val

	// change pointer ref to temp
	array = tempArray
	return array
}

// insert value at arbitrary position in array
func insert(array []int8, val int8, pos int) []int8 {
	var length = len(array)
	var tempArray = make([]int8, length+1)
	fmt.Printf("\n")

	// copy each value from start to position
	// leave the pos we want to fill empty and copy each value after that
	// eg at pos 3: 1 2 3 x 4 5 6 7 ->  1 2 3 21 4 5 6 7

	for i := 0; i < length; i++ {
		if i < pos {
			tempArray[i] = array[i]
		} else {
			tempArray[i+1] = array[i]
		}
	}

	tempArray[pos] = val
	return tempArray
}

// very similar to insert at x_pos, but we do not add a value to the empty pointer location, we move vals up
func delete(array []int8, pos int) []int8 {
	var length = len(array)
	var tempArray = make([]int8, length+1)
	fmt.Printf("\n")
	for i := 0; i < length; i++ {
		if i < pos {
			tempArray[i] = array[i]
		} else {
			tempArray[i-1] = array[i]
		}
	}
	return tempArray

}

func linearSearch(array []int8, val int8) {
	var length = len(array)
	var found = false

	for i := 0; i < length; i++ {
		if array[i] == val {
			found = true
			fmt.Printf("\nFound %d at position %d", val, i)
			break
		}
	}

	if !found {
		fmt.Printf("\n%d was not found in the list", val)
	}

}

func main() {
	var scores = []int8{90, 70, 50, 80, 60, 85}

	// length variable

	var length = len(scores)

	// for loop printing elements in array
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
	fmt.Printf("\nSuccess!\n")

	scores = append(scores, 75)

	for i := 0; i < length+1; i++ {
		fmt.Printf("%d ", scores[i])
	}

	fmt.Printf("\nSuccess!\n")

	length = len(scores)
	scores = insert(scores, 32, 3)

	for i := 0; i < length+1; i++ {
		fmt.Printf("%d ", scores[i])
	}

	fmt.Printf("\nSuccess!\n")

	scores = delete(scores, 2)

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}

	fmt.Printf("\nSuccess!\n")

	linearSearch(scores, 70)

	fmt.Printf("\nSuccess!\n")

}
