package main

import "fmt"

func main() {
	// Add a fmt.Println() statement
	// that prints 2235 * 1231
	var var1 int32
	var var2 int32
	var1 = 2235
	var2 = 1231
	var var3 int32
	var3 = var1 * var2
	fmt.Println(var3)

	fmt.Println(2235 * 1231)

	// Create the constant earthsGravity
	// here and set it to 9.80665
	const earthsGravity = 9.80665
	// Here's where we print out the gravity:
	fmt.Println(earthsGravity)

	var jellybeanCounter int8

	// Go will raise errors both for
	// jellybeanCounter being unused
	// and for "fmt" being imported
	// and unused.

	// Uncomment the following line
	// and watch the program run
	// successfully!

	fmt.Println(jellybeanCounter)

	var numOfFlavors int
	// Assign a value for numOfFlavors below:

	numOfFlavors = 57

	fmt.Println(numOfFlavors)

	// Declare flavorScale below:
	var flavorScale float32 = 5.8

	fmt.Println(flavorScale)

	
}
