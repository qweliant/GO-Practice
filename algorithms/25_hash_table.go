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

// create a slice for the hash table of type node and length CAPACITY 
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
	var length = len(key)

	for i := 0; i < length; i++{
		num += int(key[i])
	}
	
	//get the aquare in the middle
	var avg = num * int((math.Pow(5.0, 0.5) - 1)) / 2
	var numeric = avg - int(math.Floor(float64(avg)))

	return int(math.Floor(float64(numeric * CAPACITY)))
}

func put(key string, value string){
	var hash = hash(key)
	var newNode *Node = new(Node)

	newNode.key = key
	newNode.value = value
	newNode.hash = hash
	newNode.next = nil

	var node = table[hash]
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
	newNode.next = table[hash]
	table[hash] = newNode
	size++
}

func get(key string) string{
	if key == ""{
		return ""
	}
	var hash = hash(key)
	var node = table[hash]
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
	put("david", "Nigga is a murderer")
	put("grace", "Damali's friend")

	fmt.Printf("david => %s \n", get("david"))
	fmt.Printf("grace => %s \n", get("grace"))
}