package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

var leaf *Node = new(Node)

func newNode(value int) *Node {
	var node *Node = &Node{data: value, left: nil, right: nil}
	return node
}

func insert(value int, node *Node) {
	if leaf == nil {
		leaf = &Node{data: value, left: nil, right: nil}
		return
	}

	var comp = value - node.data

	// * Go to the right
	if comp < 0 {
		if node.left == nil {
			// * cretat a new node
			node.left = newNode(value)
		} else {
			insert(value, node.left)
		}
		// * Go to the left
	} else if comp > 0 {
		if node.right == nil {
			// * create a new node
			node.right = newNode(value)
		} else {
			insert(value, node.right)
		}
	}
}

// traversal in order
func inOrder(leaf *Node) {

	if leaf == nil {
		return
	}
	inOrder(leaf.left) // left to right, and is really dft
	fmt.Printf("%d\n", leaf.data)
	inOrder(leaf.right)
}

func postOrder(leaf *Node) {

	if leaf == nil {
		return
	}
	postOrder(leaf.right)
	postOrder(leaf.left)
	fmt.Printf("%d\n", leaf.data)
}

func min(node *Node) *Node {
	if node == nil || node.data == 0 {
		return nil
	}
	if node.left == nil {
		return node
	}
	return min(node.left)
}

func max(node *Node) *Node {
	if node == nil || node.data == 0 {
		return nil
	}
	if node.right == nil {
		return node
	}
	return max(node.right)
}

// func delete(value int, node *Node) *Node{
// 	if node == nil{
// 		return node
// 	}

// 	var comp = value -
// }

func main() {
	insert(60, leaf)
	insert(40, leaf)
	insert(20, leaf)
	insert(10, leaf)
	insert(30, leaf)
	insert(50, leaf)
	insert(80, leaf)
	insert(70, leaf)
	insert(90, leaf)

	fmt.Print("In order printing \n")
	inOrder(leaf)

	// fmt.Print("\n")
	// postOrder(leaf)

	// var Min = max(leaf)
	// var Max = min(leaf)

	// fmt.Print("\nMin ", Min)
	// fmt.Print("Max ", Max)

}
