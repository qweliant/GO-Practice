import "fmt"

func shellSort(array []int32, length int) {

	// Ciura gap sequence
	var gaps = []int32{701, 301, 132, 57, 23, 10, 4, 1}

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

	var scores = []int32{9, 6, 5, 8, 0, 7, 4, 3, 1, 2, 210}
	var length = len(scores)

	shellSort(scores, length)

	// for loop printing elements in array
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}
