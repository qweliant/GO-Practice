package main
import (
	"fmt"
	"math"
	"strings"
)

// node for holding
type Node struct {
	key string
	value string
	hash int
	next *Node
}

// size of the HT is declared, making it easy to adjust
const CAPACITY = 16

// create a slice for the hash table of empty nodes and length CAPACITY 
var table = make([]*Node, CAPACITY)
var size int

// checks if the table is empty
func isEmpty() bool{
	if size == 0 {
		return true
	} else {
		return false
	}
}

// hashes the key for insertion into table
func hash(key string) int{
	var num = 0
	// get the lenght of the key
	var length = len(key)

	// add the ascii character value to creat a sum 
	for i := 0; i < length; i++{

		num += int(key[i])
	}
	
	// square in the middle hash method
	var avg = num * int((math.Pow(5.0, 0.5) - 1)) / 2
	var numeric = avg - int(math.Floor(float64(avg)))


	// hash value to place into the table slice between -1 and CAPACITY - 1
	return int(math.Floor(float64(numeric * CAPACITY)))
}

func put(key string, value string){

	/*
	60 -66: 
	get index from hash function 
	create a new node, then set all the values
	*/
	var hash = hash(key)
	var newNode *Node = new(Node)

	newNode.key = key
	newNode.value = value
	newNode.hash = hash
	newNode.next = nil

	/*
	lines 80-92
	get the empty node from position at hash value

	# THIS PART HANDLES COLLISIONS
	if the node is nil break the for loop
	if the keys are 0 set the value and return
	set the node to next

	this will iterate through node pointers at a table position until we reach the end

	*/
	var node = table[hash]
	fmt.Print("Node:\n", node)

	for {
		if node == nil{
			break
		}
		if strings.Compare(node.key, key) == 0{
			node.value = value
			return
		}
		node = node.next
	}

	// set the next node to be the current hastable node
	// i.e at pos 4 : value: A
	// now is pos 4 :  value: A -> value: b
	newNode.next = table[hash]
	table[hash] = newNode
	size++
}

func get(key string) string{
	//empty keys can be checked easily
	if key == ""{
		return ""
	}

	// get the hash value of the key
	var hash = hash(key)

	// get the node from the table
	var node = table[hash]

	// if the node is nil return
	// if the node is at 0 return it
	// else get the last value enterd in the table at position hash
	for {
		if node == nil {
			break
		}
		if strings.Compare(node.key, key) == 0 {
			return node.value
		}
		node = node.next
	}
	return ""
}

func main(){
	put("david", "King is a murderer")
	put("grace", "Damali's friend")
	put("grace", "Shana's friend")

	fmt.Printf("david => %s \n", get("david"))
	fmt.Printf("grace => %s \n", get("grace"))
}