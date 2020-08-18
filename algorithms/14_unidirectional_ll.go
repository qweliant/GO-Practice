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
	// * O(n)
	// var node = head
	// for {
	// 	if node.next == nil {
	// 		var newNode *Node = &Node{data: value, next: nil}
	// 		tail = newNode
	// 		tail.next = nil
	// 		node.next = tail
	// 		break
	// 	}

	// 	node = node.next
	// }

	// * O(1)
	var newNode *Node = &Node{data: value, next: nil}
	tail.next = newNode
	tail = newNode

}

func at(value string, pos int) {
	// * O(n)
	var node = head
	var prev = node
	var i = 0
	for {
		if node.next == nil || i > pos-1 {
			break
		}
		prev = node
		node = node.next
		i++
	}

	var newNode *Node = &Node{data: value, next: node.next}
	//swap nodes
	//insert
	prev.next = newNode

}

func main() {
	initialize()
	// traversal()
	add("San Diego")
	add("Walnut")

	// traversal()
	at("Victoria", 2)
	traversal()

}
