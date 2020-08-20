package main

import "fmt"

type Node struct {
	data string
	next *Node
	prev *Node
}

var head *Node = nil
var tail *Node = new(Node)
var size int

func push(value string) {

	// init head if it is nil.
	// make sure to set tail equal to node

	if head == nil {
		head = new(Node)
		head.data = value
		tail = head
	} else {
		// links are switched for readibility and consistency of FIFO queue structure
		// make a nerw node
		// set the next to tail
		// set tail .prev to
		var newNode *Node = &Node{data: value, next: tail, prev: nil}
		tail.prev = newNode
		tail = newNode
	}
	size++
}

func pop() *Node {
	var node = head

	if node == nil {
		return nil
	}

	head = head.prev
	node.next = nil
	node.prev = nil
	size--
	return node
}

// pops each node off LL
func output() {
	fmt.Printf("Front")
	var node *Node = nil
	for {
		node = pop()
		if node == nil {
			break
		}
		fmt.Printf(" %s <- ", node.data)
	}
	fmt.Printf("Back\n")
	fmt.Println("----------------------------------------------------------")

}

func main() {
	push("1")
	push("2")
	push("3")
	push("4")
	push("5")

	output()

	push("1")
	push("2")

	output()

}
