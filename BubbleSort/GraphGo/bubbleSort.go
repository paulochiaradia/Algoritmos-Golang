// Algorítimo de bubbleSort usango Golang
// Tempo de execução 603.3 µs - i5-1135G7

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	numeros := []int{1, 3, 5, 4, 2, 6, 9, 8, 7, 0, 10}
	tamanho := len(numeros)
	bubbleSort(numeros, tamanho)

}

// Funcao para calcular o tempo de execusão
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func bubbleSort(numeros []int, tamanho int) {
	defer timeTrack(time.Now(), "BubbleSort")
	for i := 0; i < tamanho-1; i++ {
		for j := 0; j < tamanho-i-1; j++ {
			if numeros[j] > numeros[j+1] {
				temp := numeros[j]
				numeros[j] = numeros[j+1]
				numeros[j+1] = temp
			}

		}
	}
	fmt.Println(numeros)
}
