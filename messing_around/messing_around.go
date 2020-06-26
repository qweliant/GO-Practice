package main

import (
	"fmt"
	t "time"
)

func main() {

	// initialize array. arrays must be initialized to appropriate size
	var n [10]int
	var i, j int
	/* initialize elements of array n to 0.......no parenthesis!!!!!!!!!!!*/
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* set element at location i to i + 100 */
	}

	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	// print time
	fmt.Println(t.Now())

	

}
