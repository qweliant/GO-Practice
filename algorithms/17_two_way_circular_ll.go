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

func main() {
	initial()
	output()
}
