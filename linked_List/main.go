package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	list := List{}
	Paulo := Pessoa{
		Nome:  "Paulo",
		Idade: 21,
	}
	Aline := Pessoa{
		Nome:  "Aline",
		Idade: 21,
	}
	Joao := Pessoa{
		Nome:  "Joao",
		Idade: 30,
	}
	Maria := Pessoa{
		Nome:  "Maria",
		Idade: 32,
	}
	list.Append(Paulo)
	list.Append(Aline)
	list.Append(Joao)
	list.Append(Maria)
	list.Display()
	fmt.Println("--------------")
	list.Delete("Eduardo")
	list.Display()
	fmt.Println("--------------")
	list.Delete("Aline")
	list.Display()

}
