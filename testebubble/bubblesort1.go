// Algorítimo clássico de bubbleSort usango Golang
// Tempo de execução 553.1 µs - i5 9600K
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

func bubbleSort(numero []int) {
	defer timeTrack(time.Now(), "BubbleSort")
	for i := 0; i < len(numero)-1; i++ {
		for j := 0; j < len(numero)-1; j++ {
			if numero[i] < numero[j] {
				temp := numero[i]
				numero[i] = numero[j]
				numero[j] = temp
			}
		}
	}
	fmt.Println(numero)
}
