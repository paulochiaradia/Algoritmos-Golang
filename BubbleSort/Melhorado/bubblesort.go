// Algorítimo mais rápido de bubbleSort usango Golang
// Tempo de execução 517.6 µs - i5 9600K
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var numeros [11]int = [11]int{1, 3, 5, 4, 2, 6, 9, 8, 7, 0, 10}
	fmt.Println("Algoritimo BubbleSort")
	bubbleSorter(numeros)
}

// Funcao para calcular o tempo de execusão
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func bubbleSorter(numeros [11]int) {
	defer timeTrack(time.Now(), "BubbleSort")
	num := len(numeros)
	foiTrocado := true
	for foiTrocado {
		foiTrocado = false
		for i := 1; i < num; i++ {
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
