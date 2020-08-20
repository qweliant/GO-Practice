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
		fmt.Println("Departing station ->", node.data)

		node = node.next
	}
	fmt.Print("\n")

	node = tail
	for {

		if node.prev == nil {
			fmt.Println("Arriving at the first station ->", node.data)

			break
		}
		fmt.Println("Departing station ->", node.data)

		node = node.prev
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

func at(value string, pos int) {
	// * O(n)
	var node = head
	var i = 0
	for {
		if pos == 0 {
			fmt.Println("Adding new head:", head)
			var newNode *Node = &Node{data: value, next: head.next, prev: nil}
			head.next.prev = newNode
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

	var newNode *Node = &Node{data: value, next: node, prev: node.prev}
	//swap nodes
	//insert
	node.prev.next = newNode
	node.prev = newNode
}

func del(pos int) {
	// * O(n)
	var node = head
	var i = 0

	for {
		if pos == 0 {
			fmt.Println("Removing head:", head)
			head = node.next
			head.prev = nil
			return
		}

		if node.next == nil || i > pos-1 {
			break
		}
		// prev = node
		node = node.next
		i++
	}

	//swap nodes
	//insert
	// prev.next = node.next
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil
}

func main() {
	initialize()
	traversal()
	fmt.Print("-----------------------------------\n")

	addEnd("Walnut")
	traversal()
	fmt.Print("-----------------------------------\n")

	at("Victoria", 2)
	traversal()
	fmt.Print("-----------------------------------\n")

	del(0)
	fmt.Print("\n")
	traversal()
	fmt.Print("-----------------------------------\n")

	at("San Domingo", 0)
	traversal()
	fmt.Print("-----------------------------------\n")
}
