package main

import "fmt"

type Node struct {
	data string
	next *Node
	prev *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "A"
	head.next = nil
	head.prev = nil

	var nodeB *Node = &Node{data: "B", next: nil, prev: head}
	head.next = nodeB

	var nodeC *Node = &Node{data: "C", next: nil, prev: nodeB}
	nodeB.next = nodeC

	tail.data = "D"
	tail.next = head
	tail.prev = nodeC
	nodeC.next = tail
	head.prev = tail

}

func output() {
	var point = head
	for {
		fmt.Printf(" %s ->", point.data)
		point = point.next
		if point == head {
			break
		}
	}
	fmt.Printf(" Back at head %s\n\n", point.data)

	point = tail
	for {
		fmt.Printf(" %s ->", point.data)
		point = point.prev
		if point == tail {
			break
		}
	}
	fmt.Printf(" Ended at tail %s\n\n", point.data)
}

func at(value string, pos int) {
	// * O(n)
	var node = head
	var i = 0
	for {
		if pos == 0 {
			fmt.Println("Adding new head:", head)
			var newNode *Node = &Node{data: value, next: head, prev: tail}
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

	var newNode *Node = &Node{data: value, next: node.next, prev: node}
	//swap nodes
	//insert
	node.next = newNode
	newNode.next.prev = newNode
}

func delete(pos int) {
	// * O(n)
	var node = head

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
	
	var temp = node.next
	node.next = node.next.next
	node.next.prev = node
	temp.next= nil
	temp.prev = nil

}

func main() {
	initial()
	at("H", 0)
	at("E", 2)
	output()
	delete(2)
	output()

}
