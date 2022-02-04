package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"joão":   123456.40,
		"maria":  15.55,
		"junior": 1200.00,
	}

	funcsESalarios["rafael"] = 147.53
	delete(funcsESalarios, "inexistente")

	for nome, salarios := range funcsESalarios {
		fmt.Println(nome, salarios)
	}
}
