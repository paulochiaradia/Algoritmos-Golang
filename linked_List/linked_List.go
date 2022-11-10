package main

import "fmt"

// List -Struct do tipo Lista
type List struct {
	Head *Node
	Tail *Node
}

// Node - Struct do Tipo Nó
type Node struct {
	Value Pessoa
	Next  *Node
}

func (l *List) Append(pessoa Pessoa) {
	node := &Node{Value: pessoa}
	if l.Head == nil {
		l.Head = node
	}
	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
}

func (l *List) Display() {
	node := l.Head

	for node != nil {
		fmt.Printf("%s - %d\n", node.Value.Nome, node.Value.Idade)
		node = node.Next
	}
}
func (l *List) Search(nome string) int {
	node := l.Head
	for node != nil {
		if node.Value.Nome == nome {
			return 1
		}
		node = node.Next
	}
	return 0
}

func (l *List) Delete(nome string) {
	if l.Search(nome) != 0 {
		if l.Head.Value.Nome == nome {
			l.Head = l.Head.Next
			return
		}
		prev := l.Head
		node := l.Head.Next

		for node != nil {
			if node.Value.Nome == nome {
				prev.Next = node.Next
				//proximo do anterior aponta para o meu proximo
			}
			prev = node
			node = node.Next
		}
		if l.Tail == node {
			l.Tail = prev
		}
	}

}
