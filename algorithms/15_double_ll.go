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
		
		if node == nil {
			break
		}
		fmt.Println("Departing station ->", node.prev.data)
		fmt.Println("Arriving at station ->", node.data)
		fmt.Println("Next stop is station ->", node.next.data)
		fmt.Print("\n")

		node = node.next
	}
}

// func add(value string) {

// }

// func at(value string, pos int) {

// }

// func del(pos int) {

// }

func main() {
	initialize()
	traversal()
	fmt.Print("\n")

	add("San Diego")
	add("Walnut")
	traversal()
	fmt.Print("\n")

	at("Victoria", 2)
	traversal()
	fmt.Print("\n")

	del(2)
	fmt.Print("\n")
	traversal()
}
