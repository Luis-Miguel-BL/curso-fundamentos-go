package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("turbo")
}
func (b bmw) fazerBaliza() {
	fmt.Println("baliza")
}

func main() {
	var b esportivoLuxuoso = bmw{}
	b.fazerBaliza()
	b.ligarTurbo()
}
