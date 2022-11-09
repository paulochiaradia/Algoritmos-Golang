package main

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func main() {
	list := List{}

	Paulo := Pessoa{Nome: "Paulo", Sobrenome: "Neto", Idade: 21}
	Aline := Pessoa{Nome: "Aline", Sobrenome: "Souza", Idade: 21}

	list.Append(Paulo)
	list.Append(Aline)

	list.Display()
}
