package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)


// What is the time complexity for insertion into a linked list ? O(1)
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

		fmt.Println("Arriving at station ->", node.data)
		node = node.next
	}
}

func add(value string) {
	// * O(n) because LL is of size n

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

	// hera an ll, add somethin to it vs look at all values

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

func del(pos int) {
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

	//swap nodes
	//insert
	prev.next = node.next
	node.next = nil
}

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

	del(0)
	fmt.Print("\n")
	traversal()


	fmt.Println("O(1)",head.data)
	traversal()

}
