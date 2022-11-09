package main

import "fmt"

type List struct {
	Head *Node
	Tail *Node
}

type Node struct {
	value Pessoa
	Next  *Node
}

func (l *List) emptyHead() bool {
	return l.Head == nil
}
func (l *List) emptyTail() bool {
	return l.Tail == nil
}
func (l *List) Append(pessoa Pessoa) {
	node := Node{value: pessoa}
	if l.emptyHead() {
		l.Head = &node
	}
	if l.emptyTail() {
		l.Head.Next = &node
	}
	l.Tail = &node
}

func (l *List) Display() {
	node := l.Head
	for node != nil {
		fmt.Println(node.value.Nome)
		node = node.Next
	}
}
