package main

import (
	"fmt"
)

func main() {
	var numeros [11]int = [11]int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("Algoritimo BubbleSort")
	bubbleSorter(numeros)
}
func bubbleSorter(numeros [11]int) {
	num := 11
	foiTrocado := true
	for foiTrocado {
		foiTrocado = false
		var i int
		for i = 1; i < num; i++ {
			if numeros[i-1] > numeros[i] {
				var temp = numeros[i]
				numeros[i] = numeros[i-1]
				numeros[i-1] = temp
				foiTrocado = true
			}
		}
	}
	fmt.Println(numeros)
}
