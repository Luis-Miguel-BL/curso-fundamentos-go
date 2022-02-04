package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.1, 1.7, 9.9
	fmt.Println("ultima nota", notas[:0])
	fmt.Println(notas)

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("média %.2f \n", media)

}
