
import "fmt"

// Today I will be doing Dichotomous Bin search

func binSearch(array []int, val int) int {

	var length = len(array)
	var low = 0
	var high = length - 1
	var mid = (high + low) / 2
	fmt.Printf("Indices %d %d %d ", low, mid, high)

	for {
		if array[mid] == val {
			return mid
		} else if array[mid] > val {
			high = mid - 1
			mid = (high + low) / 2

		} else if array[mid] < val {
			low = mid + 1
			mid = (high + low) / 2

		} else if low >= high {
			fmt.Printf("Not Found")
			return -1
		}
	}

}

func main() {
	var scores = []int{10, 20, 30, 40, 50, 60, 70, 80, 95, 210, 590}

	// length variable
	var res int
	res = binSearch(scores, 80)
	fmt.Printf("Result: %d", res)

}
