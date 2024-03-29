package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println("Conceito = ", notaParaConceito(9.2))
	fmt.Println("Conceito = ", notaParaConceito(8))
	fmt.Println("Conceito = ", notaParaConceito(6.5))
	fmt.Println("Conceito = ", notaParaConceito(4))
	fmt.Println("Conceito = ", notaParaConceito(2.1))
}
