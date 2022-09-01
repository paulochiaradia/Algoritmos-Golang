// Algoritmo de bubbleSort usando Golang
// Tempo de execução 510.7 µs - i5-1135G7
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var numeros [12]int = [12]int{1, 3, 5, 4, 2, 6, 9, 8, 7, 0, 10}
	fmt.Println("Algoritimo BubbleSort")
	bubbleSort(numeros)
}

// Funcao para calcular o tempo de execusão
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func bubbleSort(numeros [12]int) {
	defer timeTrack(time.Now(), "BubbleSort")
	num := len(numeros)
	foiTrocado := true
	for foiTrocado {
		foiTrocado = false
		for i := 1; i < num; i++ {
			if numeros[i-1] > numeros[i] {
				temp := numeros[i]
				numeros[i] = numeros[i-1]
				numeros[i-1] = temp
				foiTrocado = true
			}
		}
	}
	fmt.Println(numeros)
}
