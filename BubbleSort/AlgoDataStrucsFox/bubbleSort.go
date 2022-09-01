// Algorítimo de bubbleSort usango Golang
// Tempo de execução 496.8 µs - i5-1135G7

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	numeros := []int{1, 3, 5, 4, 2, 6, 9, 8, 7, 0, 10}
	bubbleSort(numeros)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func bubbleSort(numeros []int) {
	defer timeTrack(time.Now(), "BubbleSort")
	for j := len(numeros) - 1; 0 < j; j-- {
		for i := 0; i < j; i++ {
			if numeros[i+1] < numeros[i] {
				numeros[i], numeros[i+1] = numeros[i+1], numeros[i]
			}
		}
	}
	fmt.Println(numeros)
}
