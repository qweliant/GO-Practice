package main
import "fmt"

/*
! https://www.geeksforgeeks.org/heap-sort/#:~:text=Heap%20sort%20is%20a%20comparison,process%20for%20the%20remaining%20elements.

A Binary Heap is a Complete Binary Tree where items are stored in
a special order such that value in a parent node is greater(or smaller)
than the values in its two children nodes. The former is called as
max heap and the latter is called min-heap. The heap can be represented
by a binary tree or array.

How to build the heap?
Heapify procedure can be applied to a node only if its children nodes are
heapified. So the heapification must be performed in the bottom-up order.

this means the bottom most parent has the heapify done first


*/

func adjustHeap(arr []int, index int, maxLen int) {
	var noLeafValue = arr[index]
	for i := 2*index + 1; i <= maxLen; i = index*2 + 1 {
		if i < maxLen && arr[i] < arr[i+1] {
			i++
		}

		if noLeafValue >= arr[i] {
			break
		}

		arr[index] = arr[i]
		index = i
	}
	arr[index] = noLeafValue
}

// build the heap
func makeHeap(arr []int, len int) {
	for i := (len - 1) / 2; i >= 0; i-- {
		adjustHeap(arr, i, len-1)
	}
}

func heapSort(arr []int, len int) {
	for i := len - 1; i > 0; i-- {
		var temp = arr[0]
		arr[0] = arr[i]
		arr[i] = temp
		adjustHeap(arr, 0, i-1)
	}
}

func main() {
	var values = []int{10, 90, 20, 80, 30, 70, 40, 60, 50}
	var length = len(values)

	fmt.Printf("Before the heap is made")
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", values[i])
	}
	fmt.Printf("\n\n")

	fmt.Printf("After the heap is made")
	makeHeap(values, length)
	for i := 0; i < length; i++{
		fmt.Printf("%d", values[i])
	}
	fmt.Printf("\n\n")

	fmt.Printf("After the heap is sorted")
	heapSort(values, length)
	for i := 0; i < length; i++{
		fmt.Printf("%d", values[i])
	}


}
