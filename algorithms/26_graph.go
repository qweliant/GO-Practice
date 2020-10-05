package main

import "fmt"

const MAX_VERTEX_SIZE = 5
const STACK_SIZE = 1000

// Vertex has value and keeps track of whether it was traversed or not
type Vertex struct {
	data    string
	visited bool
}

var top = -1 // stack save vertices

// create a slice for the number of vertices of empty int and length STACK_SIZE
var stacks = make([]int, STACK_SIZE)

// pushes element at position top
func push(element int) {
	top++
	stacks[top] = element
}

// return element
func pop() int {
	if top == -1 {
		return -1
	}
	var data = stacks[top]
	top--
	return data
}

func peek() int {
	if top == -1 {
		return -1
	}
	var data = stacks[top]
	return data
}

func isEmpty() bool {
	if top <= -1 {
		return true
	}
	return false
}

var size = 0
var vertices = make([]Vertex, MAX_VERTEX_SIZE)
var adjMatrix [MAX_VERTEX_SIZE][MAX_VERTEX_SIZE]int

func addVertex(data string) {
	var vertex Vertex
	vertex.data = data
	vertex.visited = false
	vertices[size] = vertex
	size++
}

func addEdge(from int, to int) { // add adjacenct edges
	adjMatrix[from][to] = 1 // A -> B ! B -> A
}

func clear() {
	for i := 0; i < size; i++ {
		vertices[i].visited = false
	}
}

func depthFirstSearch() {
	vertices[0].visited = true
	fmt.Printf("%s", vertices[0].data)
	push(0)
	for {
		if isEmpty() {
			break
		}
		var row = peek()

		// get adjacency vertices not visited
		var col = findUnvisistedVertex(row)
		if col == -1 {
			pop()
		} else {
			vertices[col].visited = true
			fmt.Printf("->%s", vertices[col].data)
			push(col)
		}
	}
	clear()
}

func findUnvisistedVertex(row int) int {
	for col := 0; col < size; col++ {
		if adjMatrix[row][col] == 1 && !vertices[col].visited {
			return col
		}
	}
	return -1
}

func printGraph() {
	fmt.Printf("2D array traversal\n")
	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf(" %s", vertices[i].data)
	}
	fmt.Printf("\n")

	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s ", vertices[i].data)
		for j := 0; j < MAX_VERTEX_SIZE; j++ {
			fmt.Printf("%d ", adjMatrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	addVertex("A")
	addVertex("B")
	addVertex("C")
	addVertex("D")
	addVertex("E")

	addEdge(0, 1)
	addEdge(0, 2)
	addEdge(0, 3)
	addEdge(1, 2)
	addEdge(1, 3)
	addEdge(2, 3)
	addEdge(3, 4)

	// ? Two Dim array traversal output
	printGraph()
	fmt.Print("\nDFS Traversal:\n")
	depthFirstSearch()
}
