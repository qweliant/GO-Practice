package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initialize() {

	head.data = "San Francisco"
	head.next = nil

	var Oakland *Node = &Node{data: "Oakland", next: nil}
	head.next = Oakland

	var Berkely *Node = &Node{data: "Berkely", next: nil}
	Oakland.next = Berkely

	var Fremont *Node = &Node{data: "Fremont", next: nil}
	Berkely.next = Fremont

	// remember this pattern for linked list
	// make the last insertion, tail = last value
	tail = Fremont
	tail.next = nil
	Berkely.next = tail
}

func traversal() {

	var node = head
	for {
		if node == nil {
			break
		}

		fmt.Println("Arriving at station", node.data)
		node = node.next
	}
}

func add(value string) {

	var node = head
	for {
		if node.next == nil {
			var newNode *Node = &Node{data: value, next: nil}
			tail = newNode
			tail.next = nil
			node.next = tail
			break
		}

		fmt.Println("Arriving at station", node.data)
		node = node.next
	}

}
func main() {
	initialize()
	// traversal()
	add("San Diego")
	traversal()

}
