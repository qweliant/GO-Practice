package main

import "fmt"

type Node struct {
	data string
	next *Node
	prev *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initialize() {

	head.data = "San Francisco"
	head.next = nil

	var Oakland *Node = &Node{data: "Oakland", next: nil, prev: nil}
	head.next = Oakland
	Oakland.prev = head

	var Berkely *Node = &Node{data: "Berkely", next: nil, prev: nil}
	Oakland.next = Berkely
	Berkely.prev = Oakland

	var Fremont *Node = &Node{data: "Fremont", next: nil, prev: nil}
	Berkely.next = Fremont
	Fremont.prev = Berkely

	// remember this pattern for linked list
	// make the last insertion, tail = last value
	tail = Fremont
	tail.next = nil
	tail.prev = Berkely
	Berkely.next = tail
}

func traversal() {

	var node = head
	for {

		if node.next == nil {
			fmt.Println("Arriving at last station ->", node.data)

			break
		}
		fmt.Println("Arriving at station ->", node.data)

		node = node.next
	}
	fmt.Print("\n")
	
}

func addEnd(value string) {

	// prev point to info that was in tail
	var newNode *Node = &Node{data: value, next: nil, prev: tail}
	// tail.next is now newNode
	tail.next = newNode
	// new node is now tail
	tail = newNode
}

// func at(value string, pos int) {
// 	for {

// 		if node.prev == nil {
// 			fmt.Println("Arriving at initial station ->", node.data)

// 			break
// 		}
// 		fmt.Println("Arriving at station ->", node.data)

// 		node = node.prev
// 	}
// 	fmt.Print("\n")
// }

// func del(pos int) {

// }

func main() {
	initialize()
	traversal()

	addEnd("Walnut")
	traversal()

	// at("Victoria", 2)
	// traversal()
	// fmt.Print("\n")

	// del(2)
	// fmt.Print("\n")
	// traversal()
}
