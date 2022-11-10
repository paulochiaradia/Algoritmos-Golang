package main

import "fmt"

func main() {
	stack := Stack{}

	stack.Push("Paulo")
	stack.Push("Aline")
	stack.Push("Pedro")

	fmt.Println(stack.Pop()) //Pedro
	fmt.Println(stack.Pop()) //Aline
	fmt.Println(stack.Pop()) //Paulo
}
