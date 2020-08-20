
import "fmt"


func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	// length variable
	var length = len(scores)

	sort(scores, length)

	// for loop printing elements in array
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func sort(arr []int, length int) {
	
	var minIndex int
	
	for i := 0; i < length-1; i++ {
		
		minIndex = i
		var minValue = arr[minIndex]
		
		for j := i; j < length-1; j++ {
			if minValue > arr[j+1] {
				minValue = arr[j+1]
				minIndex = j + 1
			}
		}
		
		if i != minIndex {
			var temp = arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
	}
}


