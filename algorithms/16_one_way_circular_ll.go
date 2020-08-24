package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "A"
	head.next = nil

	var nodeB *Node = &Node{data: "B", next: nil}
	head.next = nodeB

	var nodeC *Node = &Node{data: "C", next: nil}
	nodeB.next = nodeC

	tail.data = "D"
	tail.next = head
	nodeC.next = tail

}

func at(value string, pos int) {
	// * O(n)
	var node = head
	var i = 0
	for {
		if pos == 0 {
			fmt.Println("Adding new head:", head)
			var newNode *Node = &Node{data: value, next: head}
			tail.next = newNode
			head = newNode

			return
		}
		if node.next == nil || i > pos-1 {
			break
		}
		// prev = node
		node = node.next
		i++
	}

	var newNode *Node = &Node{data: value, next: node.next}
	//swap nodes
	//insert
	node.next = newNode
}

func delete(pos int) {
	// * O(n)
	var node = head
	var prev = node

	var i = 0
	for {
		if pos == 0 {
			fmt.Println("Deleting head:", head)
			head = node.next
			tail.next = head

			return
		}
		if node.next == nil || i > pos-1 {
			break
		}

		node = node.next
		i++
	}

	//swap nodes
	//insert
	prev.next = node.next
	node.next = nil
}

func output(node *Node) {
	var point = node
	for {
		fmt.Printf(" %s ->", point.data)
		point = point.next
		if point == head {
			break
		}
	}
	fmt.Printf(" Back at %s\n\n", point.data)
}

func main() {
	initial()

	fmt.Printf("Insert new node E\n")
	at("E", 2)
	output(head)
	at("H", 0)
	output(head)
	delete(0)
	output(head)

}
