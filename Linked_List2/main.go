package main

import "fmt"

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToHead(3)
	linkedList.AddToHead(4)
	fmt.Println(linkedList.headNode.property)
	linkedList.IterateList()
	linkedList.AddToEnd(9)
	node := linkedList.LastNode()
	fmt.Println(node.property)
	fmt.Println("----------------------")
	linkedList.IterateList()
	node1 := linkedList.NodeWithValue(3)
	fmt.Println(node1.property)
	fmt.Println("----------------------")
	linkedList.AddAfter(
		1,
		2,
	)
	linkedList.IterateList()

}
