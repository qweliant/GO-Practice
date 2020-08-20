package main
import "fmt"
type Node struct {
	data string
	next *Node
}

var head *Node = nil
var size int

func push(value string){
	if head == nil{
		head = new(Node)
		head.data = value
	} else {
		var newNode = &Node{data: value, next: head}
		head = newNode
	}
	size++
}

func pop() *Node{
	var node = head

	if node == nil {
		return nil
	}

	head = head.next
	node.next = nil
	size--
	return node
}

func output(){
	fmt.Printf("Top")
	var node *Node = nil
	for {
		node = pop()
		if node == nil {
			break
		}
		fmt.Printf(" %s -> ", node.data)
	}
	fmt.Printf("Bottom\n")
	fmt.Println("----------------------------------------------------------")
}

func main(){
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