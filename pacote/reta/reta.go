package main

import "math"

// Ponto testee sfsfs
type Ponto struct {
	x float64
	y float64
}

// catetos sfsdfs
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia asdfsf safas asdfsdf
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
